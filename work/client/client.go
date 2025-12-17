package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpy/wechatgo/client"
	"github.com/wechatpy/wechatgo/session"
)

// WorkClient 企业微信客户端
type WorkClient struct {
	*client.BaseClient
	AppID  string
	Secret string

	// API 模块
	User    *UserAPI    `json:"-"` // 用户管理
	Dept    *DeptAPI    `json:"-"` // 部门管理
	Tag     *TagAPI     `json:"-"` // 标签管理
	Message *MessageAPI `json:"-"` // 消息管理
	Media   *MediaAPI   `json:"-"` // 媒体管理
	Contact *ContactAPI `json:"-"` // 客户联系
	OA      *OAAPI      `json:"-"` // 办公应用
}

// NewWorkClient 创建企业微信客户端
func NewWorkClient(corpID, corpSecret string, storage session.Storage) *WorkClient {
	baseClient := client.NewBaseClient(corpID, storage, APIBaseURL)

	c := &WorkClient{
		BaseClient: baseClient,
		AppID:      corpID,
		Secret:     corpSecret,
	}

	// 初始化 API 模块
	c.User = NewUserAPI(c)
	c.Dept = NewDeptAPI(c)
	c.Tag = NewTagAPI(c)
	c.Message = NewMessageAPI(c)
	c.Media = NewMediaAPI(c)
	c.Contact = NewContactAPI(c)
	c.OA = NewOAAPI(c)

	return c
}

// APIBaseURL 企业微信API基础URL
const APIBaseURL = "https://qyapi.weixin.qq.com/cgi-bin/"

// TokenURL 获取 access token 的 URL
const TokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"

// 错误定义
var (
	ErrAccessTokenNotFound    = fmt.Errorf("access_token not found in response")
	ErrAccessTokenInvalidType = fmt.Errorf("access_token is not a string")
)

// FetchAccessToken 获取 access token
func (c *WorkClient) FetchAccessToken() error {
	params := map[string]string{
		"corpid":     c.AppID,
		"corpsecret": c.Secret,
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

// accessTokenKey 获取 access token 存储键
func (c *WorkClient) accessTokenKey() string {
	return c.AppID + "_" + c.Secret[:10] + "_access_token"
}

// GetAccessToken 获取 access token（重写以支持自动刷新）
func (c *WorkClient) GetAccessToken() (string, error) {
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

// fetchToken 获取 token（内部方法，不自动添加 access_token 参数）
func (c *WorkClient) fetchToken(url string, params map[string]string) (map[string]interface{}, error) {
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
