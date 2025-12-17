package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpy/wechatgo/client/api"
	"github.com/wechatpy/wechatgo/session"
)

const (
	// APIBaseURL 微信 API 基础 URL
	APIBaseURL = "https://api.weixin.qq.com/cgi-bin/"
	// TokenURL 获取 access token 的 URL
	TokenURL = "https://api.weixin.qq.com/cgi-bin/token"
)

// Client 微信客户端
type Client struct {
	*BaseClient
	Secret string

	// API 模块
	User          *api.UserAPI
	Message       *api.MessageAPI
	Menu          *api.MenuAPI
	Media         *api.MediaAPI
	Template      *api.TemplateAPI
	QRCode        *api.QRCodeAPI
	Tag           *api.TagAPI
	CustomService *api.CustomServiceAPI
	DataCube      *api.DataCubeAPI
	Device        *api.DeviceAPI
	POI           *api.POIAPI
	WiFi          *api.WiFiAPI
	Misc          *api.MiscAPI
}

// NewClient 创建微信客户端
func NewClient(appID, secret string, storage session.Storage) *Client {
	baseClient := NewBaseClient(appID, storage, APIBaseURL)

	client := &Client{
		BaseClient: baseClient,
		Secret:     secret,
	}

	// 初始化 API 模块
	client.User = api.NewUserAPI(client)
	client.Message = api.NewMessageAPI(client)
	client.Menu = api.NewMenuAPI(client)
	client.Media = api.NewMediaAPI(client)
	client.Template = api.NewTemplateAPI(client)
	client.QRCode = api.NewQRCodeAPI(client)
	client.Tag = api.NewTagAPI(client)
	client.CustomService = api.NewCustomServiceAPI(client)
	client.DataCube = api.NewDataCubeAPI(client)
	client.Device = api.NewDeviceAPI(client)
	client.POI = api.NewPOIAPI(client)
	client.WiFi = api.NewWiFiAPI(client)
	client.Misc = api.NewMiscAPI(client)

	return client
}

// FetchAccessToken 获取 access token
func (c *Client) FetchAccessToken() error {
	params := map[string]string{
		"grant_type": "client_credential",
		"appid":      c.AppID,
		"secret":     c.Secret,
	}

	// 不使用 BaseClient.Get，因为它会自动添加 access_token 参数
	result, err := c.fetchToken(TokenURL, params)
	if err != nil {
		return err
	}

	accessTokenIf, ok := result["access_token"]
	if !ok {
		return fmt.Errorf("access_token not found in response")
	}
	accessToken, ok := accessTokenIf.(string)
	if !ok {
		return fmt.Errorf("access_token is not a string, got: %T", accessTokenIf)
	}

	expiresIn := 7200
	if exp, ok := result["expires_in"].(float64); ok {
		expiresIn = int(exp)
	}

	return c.SetAccessToken(accessToken, expiresIn)
}

// fetchToken 获取 token（内部方法，不自动添加 access_token 参数）
func (c *Client) fetchToken(url string, params map[string]string) (map[string]interface{}, error) {
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

// GetAccessToken 获取 access token（重写以支持自动刷新）
func (c *Client) GetAccessToken() (string, error) {
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
