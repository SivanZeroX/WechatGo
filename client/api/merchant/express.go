package merchant

import (
	"fmt"
)

// ExpressAPI 快递相关API
type ExpressAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewExpressAPI 创建快递API
func NewExpressAPI(client interface {
	Post(url string, data interface{}) (map[string]interface{}, error)
	Get(url string, params map[string]string) (map[string]interface{}, error)
}) *ExpressAPI {
	return &ExpressAPI{
		BaseAPI: client,
	}
}

// GetExpressTemplates 获取快递模板
func (api *ExpressAPI) GetExpressTemplates() ([]map[string]interface{}, error) {
	// TODO: 实现获取快递模板逻辑
	return nil, fmt.Errorf("not implemented")
}
