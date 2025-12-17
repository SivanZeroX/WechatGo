package merchant

import (
	"fmt"
)

// CategoryAPI 分类相关API
type CategoryAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewCategoryAPI 创建分类API
func NewCategoryAPI(client interface {
	Post(url string, data interface{}) (map[string]interface{}, error)
	Get(url string, params map[string]string) (map[string]interface{}, error)
}) *CategoryAPI {
	return &CategoryAPI{
		BaseAPI: client,
	}
}

// GetCategory 获取分类
func (api *CategoryAPI) GetCategory(categoryID string) (map[string]interface{}, error) {
	// TODO: 实现获取分类逻辑
	return nil, fmt.Errorf("not implemented")
}
