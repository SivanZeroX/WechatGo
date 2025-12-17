package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpy/wechatgo/client"
	"github.com/wechatpy/wechatgo/session"
)

// IotClient 微信IoT客户端
type IotClient struct {
	*client.BaseClient
	AppID  string
	Secret string

	// API 模块
	Cloud  *CloudAPI
	Device *DeviceAPI
}

// NewIotClient 创建IoT客户端
func NewIotClient(appID, secret string, storage session.Storage) *IotClient {
	baseClient := client.NewBaseClient(appID, storage, APIBaseURL)

	c := &IotClient{
		BaseClient: baseClient,
		AppID:      appID,
		Secret:     secret,
	}

	// 初始化 API 模块
	c.Cloud = NewCloudAPI(c)
	c.Device = NewDeviceAPI(c)

	return c
}

// APIBaseURL IoT API 基础 URL
const APIBaseURL = "https://api.weixin.qq.com/ilink/api/"

// TokenURL 获取 access token 的 URL
const TokenURL = "https://api.weixin.qq.com/cgi-bin/token"

// FetchAccessToken 获取 access token
func (c *IotClient) FetchAccessToken() error {
	params := map[string]string{
		"grant_type": "client_credential",
		"appid":      c.AppID,
		"secret":     c.Secret,
	}

	result, err := c.fetchToken(TokenURL, params)
	if err != nil {
		return err
	}

	accessTokenIf, ok := result["access_token"]
	if !ok {
		return ErrAccessTokenNotFound
	}
	accessToken, ok := accessTokenIf.(string)
	if !ok {
		return ErrAccessTokenInvalidType
	}

	expiresIn := 7200
	if exp, ok := result["expires_in"].(float64); ok {
		expiresIn = int(exp)
	}

	return c.SetAccessToken(accessToken, expiresIn)
}

// fetchToken 获取 token（内部方法，不自动添加 access_token 参数）
func (c *IotClient) fetchToken(url string, params map[string]string) (map[string]interface{}, error) {
	// 构建 URL 参数
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// 创建新的HTTP客户端
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// 检查错误
	if errcode, ok := result["errcode"]; ok {
		errcodeInt := int(errcode.(float64))
		if errcodeInt != 0 {
			errmsg := ""
			if msg, ok := result["errmsg"]; ok {
				errmsg = msg.(string)
			}
			return nil, fmt.Errorf("fetch token error: %d - %s", errcodeInt, errmsg)
		}
	}

	return result, nil
}

// accessTokenKey 获取 access token 存储键
func (c *IotClient) accessTokenKey() string {
	return c.AppID + "_" + c.Secret[:10] + "_access_token"
}

// GetAccessToken 获取 access token（重写以支持自动刷新）
func (c *IotClient) GetAccessToken() (string, error) {
	token, err := c.BaseClient.GetAccessToken()
	if err != nil {
		// Token 不存在或已过期，刷新
		if err := c.FetchAccessToken(); err != nil {
			return "", err
		}
		return c.BaseClient.GetAccessToken()
	}
	return token, nil
}
