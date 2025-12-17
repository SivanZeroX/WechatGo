package merchant

import (
	"fmt"
)

// CommonAPI 通用API
type CommonAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewCommonAPI 创建通用API
func NewCommonAPI(client interface {
	Post(url string, data interface{}) (map[string]interface{}, error)
	Get(url string, params map[string]string) (map[string]interface{}, error)
}) *CommonAPI {
	return &CommonAPI{
		BaseAPI: client,
	}
}

// GetMerchantInfo 获取商户信息
func (api *CommonAPI) GetMerchantInfo() (map[string]interface{}, error) {
	// TODO: 实现获取商户信息逻辑
	return nil, fmt.Errorf("not implemented")
}
