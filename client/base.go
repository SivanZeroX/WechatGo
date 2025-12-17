package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"

	"github.com/wechatpy/wechatgo"
	"github.com/wechatpy/wechatgo/client/api"
	"github.com/wechatpy/wechatgo/logger"
	"github.com/wechatpy/wechatgo/session"
)

const (
	defaultTimeout = 30 * time.Second
)

var (
	// 全局HTTP传输层，复用TCP连接
	defaultTransport = &http.Transport{
		MaxIdleConns:        100,              // 最大空闲连接数
		MaxIdleConnsPerHost: 10,               // 每主机最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时
		DisableCompression:  false,            // 启用压缩
	}

	// 全局HTTP客户端，复用连接池
	defaultHTTPClient = &http.Client{
		Transport: defaultTransport,
		Timeout:   defaultTimeout,
	}

	// 初始化锁
	httpClientInit sync.Once
)

// BaseClient 基础客户端
type BaseClient struct {
	AppID      string
	httpClient *http.Client
	session    session.Storage
	autoRetry  bool
	apiBaseURL string
	logger     logger.Logger

	// 性能优化：缓存常用数据
	mu               sync.RWMutex
	jsonMarshalCache map[string][]byte // JSON序列化缓存
}

// NewBaseClient 创建基础客户端
func NewBaseClient(appID string, storage session.Storage, apiBaseURL string) *BaseClient {
	if storage == nil {
		storage = session.NewMemoryStorage()
	}

	// 初始化HTTP客户端（使用全局优化配置）
	httpClientInit.Do(func() {
		// 全局客户端已配置好Transport
	})

	return &BaseClient{
		AppID:            appID,
		httpClient:       defaultHTTPClient,
		session:          storage,
		autoRetry:        true,
		apiBaseURL:       apiBaseURL,
		logger:           logger.New(),
		jsonMarshalCache: make(map[string][]byte, 100), // 缓存100个JSON序列化结果
	}
}

// WithLogger 设置logger
func (c *BaseClient) WithLogger(logger logger.Logger) *BaseClient {
	c.logger = logger
	return c
}

// accessTokenKey 获取 access token 存储键
func (c *BaseClient) accessTokenKey() string {
	return fmt.Sprintf("%s_access_token", c.AppID)
}

// expiresAtKey 获取过期时间存储键
func (c *BaseClient) expiresAtKey() string {
	return fmt.Sprintf("%s_access_token_expires_at", c.AppID)
}

// GetAccessToken 获取 access token
func (c *BaseClient) GetAccessToken() (string, error) {
	token, err := c.session.Get(c.accessTokenKey())
	if err != nil {
		return "", err
	}

	if token != "" {
		// 检查是否过期
		expiresAtStr, err := c.session.Get(c.expiresAtKey())
		if err == nil && expiresAtStr != "" {
			var expiresAt int64
			if err := json.Unmarshal([]byte(expiresAtStr), &expiresAt); err == nil {
				if time.Now().Unix() < expiresAt-60 {
					return token, nil
				}
			}
		} else {
			// 用户提供的 token，直接返回
			return token, nil
		}
	}

	// Token 不存在或已过期，需要刷新
	return "", fmt.Errorf("access token expired or not found")
}

// SetAccessToken 设置 access token
func (c *BaseClient) SetAccessToken(token string, expiresIn int) error {
	if err := c.session.Set(c.accessTokenKey(), token, time.Duration(expiresIn)*time.Second); err != nil {
		return err
	}

	expiresAt := time.Now().Unix() + int64(expiresIn)
	expiresAtData, err := json.Marshal(expiresAt)
	if err != nil {
		return fmt.Errorf("failed to marshal expiresAt: %w", err)
	}
	return c.session.Set(c.expiresAtKey(), string(expiresAtData), time.Duration(expiresIn)*time.Second)
}

// Request 发送 HTTP 请求
func (c *BaseClient) Request(method, urlOrEndpoint string, params map[string]string, data interface{}) (map[string]interface{}, error) {
	url := urlOrEndpoint
	if urlOrEndpoint[0] == '/' {
		url = c.apiBaseURL + urlOrEndpoint
	}

	// 记录请求开始
	timer := logger.StartTimer()
	c.logger.Info("开始发送HTTP请求",
		logger.String("method", method),
		logger.String("url", url),
		logger.Int("params_count", len(params)),
	)

	// 添加 access_token 参数
	if params == nil {
		params = make(map[string]string)
	}
	if _, ok := params["access_token"]; !ok {
		token, err := c.GetAccessToken()
		if err != nil {
			c.logger.Error("获取access token失败", err)
			return nil, err
		}
		params["access_token"] = token
	}

	// 构建请求
	var body io.Reader
	if data != nil {
		// 性能优化：使用JSON缓存减少重复序列化
		jsonData, err := c.marshalJSON(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// 添加查询参数
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	// 处理错误
	response, err := c.handleResult(result, method, urlOrEndpoint, params, data)
	if err != nil {
		timer(logger.Fields{"duration_field": "duration"})
		c.logger.Error("API调用失败", err,
			logger.String("method", method),
			logger.String("url", url),
		)
	} else {
		timer(logger.Fields{"duration_field": "duration"})
		c.logger.Debug("API调用成功",
			logger.String("method", method),
			logger.String("url", url),
		)
	}
	return response, err
}

// handleResult 处理响应结果
func (c *BaseClient) handleResult(result map[string]interface{}, method, url string, params map[string]string, data interface{}) (map[string]interface{}, error) {
	// 检查错误码
	if errcode, ok := result["errcode"]; ok {
		errcodeFloat, ok := errcode.(float64)
		if !ok {
			return nil, fmt.Errorf("invalid errcode type: %T", errcode)
		}
		errcodeInt := int(errcodeFloat)
		if errcodeInt != 0 {
			errmsg := ""
			if msg, ok := result["errmsg"]; ok {
				if errmsgStr, ok := msg.(string); ok {
					errmsg = errmsgStr
				}
			}

			// 自动重试 token 过期错误
			if c.autoRetry && (errcodeInt == int(wechatgo.InvalidCredential) ||
				errcodeInt == int(wechatgo.InvalidAccessToken) ||
				errcodeInt == int(wechatgo.ExpiredAccessToken)) {
				// Token 过期，清除缓存并重试
				c.logger.Warn("AccessToken过期，正在刷新",
					logger.Int("errcode", errcodeInt),
					logger.String("errmsg", errmsg),
				)
				c.session.Delete(c.accessTokenKey())
				c.session.Delete(c.expiresAtKey())
				return c.Request(method, url, params, data)
			}

			// API 频率限制
			if errcodeInt == int(wechatgo.OutOfAPIFreqLimit) {
				c.logger.Warn("API调用频率受限",
					logger.Int("errcode", errcodeInt),
					logger.String("errmsg", errmsg),
				)
				return nil, wechatgo.NewAPILimitedError(errcodeInt, errmsg, nil, nil)
			}

			c.logger.Error("API返回错误",
				fmt.Errorf("%d: %s", errcodeInt, errmsg),
				logger.Int("errcode", errcodeInt),
				logger.String("errmsg", errmsg),
			)
			return nil, wechatgo.NewClientError(errcodeInt, errmsg, nil, nil)
		}
	}

	return result, nil
}

// Get 发送 GET 请求
func (c *BaseClient) Get(url string, params map[string]string) (map[string]interface{}, error) {
	return c.Request("GET", url, params, nil)
}

// Post 发送 POST 请求
func (c *BaseClient) Post(url string, data interface{}) (map[string]interface{}, error) {
	return c.Request("POST", url, nil, data)
}

// GetRaw 发送原始 HTTP GET 请求
func (c *BaseClient) GetRaw(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

// PostRaw 发送原始 HTTP POST 请求
func (c *BaseClient) PostRaw(url string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

// HTTPClientWrapper HTTP客户端包装器
type HTTPClientWrapper struct {
	*BaseClient
}

// Get 实现HTTPClient接口
func (w *HTTPClientWrapper) Get(url string) (*http.Response, error) {
	return w.GetRaw(url)
}

// Post 实现HTTPClient接口
func (w *HTTPClientWrapper) Post(url string) (*http.Response, error) {
	return w.PostRaw(url)
}

// AsHTTPClient 将BaseClient转换为HTTPClient
func (c *BaseClient) AsHTTPClient() api.HTTPClient {
	return &HTTPClientWrapper{c}
}

// Upload 上传文件（实现API接口）
func (c *BaseClient) Upload(url, fileName string, file io.Reader) (map[string]interface{}, error) {
	// 构建完整的URL
	fullURL := url
	if url[0] == '/' {
		fullURL = c.apiBaseURL + url
	}

	// 创建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	// 添加access_token参数
	if token, err := c.GetAccessToken(); err == nil {
		if err := writer.WriteField("access_token", token); err != nil {
			return nil, fmt.Errorf("failed to write access_token: %w", err)
		}
	}

	writer.Close()

	// 创建请求
	req, err := http.NewRequest("POST", fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// 检查错误
	return c.handleResult(result, "POST", url, nil, nil)
}

// marshalJSON 带缓存的JSON序列化
func (c *BaseClient) marshalJSON(data interface{}) ([]byte, error) {
	// 对于简单数据类型，不使用缓存
	if _, ok := data.(string); ok {
		return json.Marshal(data)
	}

	// 生成缓存键
	cacheKey := fmt.Sprintf("%T:%v", data, data)

	// 尝试从缓存获取
	c.mu.RLock()
	if cached, exists := c.jsonMarshalCache[cacheKey]; exists {
		c.mu.RUnlock()
		return cached, nil
	}
	c.mu.RUnlock()

	// 序列化
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 添加到缓存（带写锁）
	c.mu.Lock()
	// 检查缓存大小，超过限制则清理
	if len(c.jsonMarshalCache) >= 100 {
		// 清理一半缓存
		clearCount := 0
		for key := range c.jsonMarshalCache {
			delete(c.jsonMarshalCache, key)
			clearCount++
			if clearCount >= 50 {
				break
			}
		}
	}
	c.jsonMarshalCache[cacheKey] = result
	c.mu.Unlock()

	return result, nil
}
