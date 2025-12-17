package api

import (
	"fmt"
)

// PartnerOrderAPI 合作伙伴订单API
type PartnerOrderAPI struct {
	*BaseAPI
}

// NewPartnerOrderAPI 创建合作伙伴订单API
func NewPartnerOrderAPI(client *Client) *PartnerOrderAPI {
	return &PartnerOrderAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// CreateOrder 创建订单
func (api *PartnerOrderAPI) CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	// TODO: 实现创建订单逻辑
	return nil, fmt.Errorf("not implemented")
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	AppID          string `json:"appid"`            // 公众号ID
	MchID          string `json:"mch_id"`           // 商户号
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TotalFee       int    `json:"total_fee"`        // 总金额（分）
	Body           string `json:"body"`             // 商品描述
	SpbillCreateIP string `json:"spbill_create_ip"` // 终端IP
	NotifyURL      string `json:"notify_url"`       // 通知地址
	TradeType      string `json:"trade_type"`       // 交易类型
}

// CreateOrderResponse 创建订单响应
type CreateOrderResponse struct {
	PrepayID string `json:"prepay_id"` // 预支付交易会话标识
}

// QueryOrder 查询订单
func (api *PartnerOrderAPI) QueryOrder(outTradeNo string) (*QueryOrderResponse, error) {
	// TODO: 实现查询订单逻辑
	return nil, fmt.Errorf("not implemented")
}

// QueryOrderResponse 查询订单响应
type QueryOrderResponse struct {
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
	TradeState string `json:"trade_state"`  // 交易状态
	TotalFee   int    `json:"total_fee"`    // 总金额（分）
}
