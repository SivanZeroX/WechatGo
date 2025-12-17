package merchant

import (
	"fmt"
)

// OrderAPI 订单相关API
type OrderAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewOrderAPI 创建订单API
func NewOrderAPI(client interface {
	Post(url string, data interface{}) (map[string]interface{}, error)
	Get(url string, params map[string]string) (map[string]interface{}, error)
}) *OrderAPI {
	return &OrderAPI{
		BaseAPI: client,
	}
}

// GetOrder 获取订单详情
func (api *OrderAPI) GetOrder(orderID string) (map[string]interface{}, error) {
	// TODO: 实现获取订单详情逻辑
	return nil, fmt.Errorf("not implemented")
}

// UpdateOrderStatus 更新订单状态
func (api *OrderAPI) UpdateOrderStatus(orderID, status string) error {
	// TODO: 实现更新订单状态逻辑
	return fmt.Errorf("not implemented")
}
