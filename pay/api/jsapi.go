package api

// JsAPIAPI JSAPI接口
type JsAPIAPI struct {
	*BaseAPI
}

// NewJsAPIAPI 创建JSAPI接口
func NewJsAPIAPI(client Client) *JsAPIAPI {
	return &JsAPIAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// GetPayParams 获取支付参数
func (api *JsAPIAPI) GetPayParams(prepayID string) (map[string]string, error) {
	return api.client.GenerateJSAPIPayParams(prepayID)
}

// GetAccessToken 获取access token（微信支付专用）
func (api *JsAPIAPI) GetAccessToken() (string, error) {
	// 这里应该调用微信支付API获取access token
	// TODO: 实现具体逻辑
	return "", nil
}
