package api

import (
	"net/http"
)

// BaseAPI V3版本基础API
type BaseAPI struct {
	Client *Client
}

// Client V3版HTTP客户端接口
type Client interface {
	HTTPClient
	GetMchID() string
	GetAPIKey() string
}

// HTTPClient HTTP客户端接口
type HTTPClient interface {
	Post(url string, data []byte, headers map[string]string) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

// NewBaseAPI 创建基础API
func NewBaseAPI(client *Client) *BaseAPI {
	return &BaseAPI{
		Client: client,
	}
}

// CommonResponse 通用响应
type CommonResponse struct {
	Code    string `json:"code"`    // 状态码
	Message string `json:"message"` // 消息
}
