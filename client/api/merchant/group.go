package merchant

import (
	"fmt"
)

// GroupAPI 分组相关API
type GroupAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewGroupAPI 创建分组API
func NewGroupAPI(client interface {
	Post(url string, data interface{}) (map[string]interface{}, error)
	Get(url string, params map[string]string) (map[string]interface{}, error)
}) *GroupAPI {
	return &GroupAPI{
		BaseAPI: client,
	}
}

// GetGroups 获取分组列表
func (api *GroupAPI) GetGroups() ([]map[string]interface{}, error) {
	// TODO: 实现获取分组列表逻辑
	return nil, fmt.Errorf("not implemented")
}
