package api

import (
	"net/http"
)

// HTTPClient HTTP客户端接口
type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url string) (*http.Response, error)
	GetRaw(url string) (*http.Response, error)
}

// BaseAPI API 基类
type BaseAPI struct {
	client interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
	httpClient HTTPClient
}

// NewBaseAPI 创建 API 基类
func NewBaseAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *BaseAPI {
	return &BaseAPI{client: client}
}

// NewBaseAPIWithClient 创建 API 基类（带HTTP客户端）
func NewBaseAPIWithClient(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}, httpClient HTTPClient) *BaseAPI {
	return &BaseAPI{
		client:     client,
		httpClient: httpClient,
	}
}

// Get 发送 GET 请求
func (api *BaseAPI) Get(url string, params map[string]string) (map[string]interface{}, error) {
	return api.client.Get(url, params)
}

// Post 发送 POST 请求
func (api *BaseAPI) Post(url string, data interface{}) (map[string]interface{}, error) {
	return api.client.Post(url, data)
}

// GetAccessToken 获取 access token
func (api *BaseAPI) GetAccessToken() (string, error) {
	return api.client.GetAccessToken()
}
