package api

import (
	"fmt"
)

// EcommerceAPI 电商相关API
type EcommerceAPI struct {
	*BaseAPI
}

// NewEcommerceAPI 创建电商API
func NewEcommerceAPI(client *Client) *EcommerceAPI {
	return &EcommerceAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// CreateSubMchID 创建子商户号
func (api *EcommerceAPI) CreateSubMchID(req *CreateSubMchIDRequest) (*CreateSubMchIDResponse, error) {
	// TODO: 实现创建子商户号逻辑
	return nil, fmt.Errorf("not implemented")
}

// CreateSubMchIDRequest 创建子商户号请求
type CreateSubMchIDRequest struct {
	SubMchID     string `json:"sub_mchid"`     // 子商户号
	AppID        string `json:"appid"`         // 公众号ID
	MerchantName string `json:"merchant_name"` // 商户名称
	ServicePhone string `json:"service_phone"` // 客服电话
}

// CreateSubMchIDResponse 创建子商户号响应
type CreateSubMchIDResponse struct {
	SubMchID string `json:"sub_mchid"` // 子商户号
}

// ApplyWithdraw 申请提现
func (api *EcommerceAPI) ApplyWithdraw(req *ApplyWithdrawRequest) (*ApplyWithdrawResponse, error) {
	// TODO: 实现申请提现逻辑
	return nil, fmt.Errorf("not implemented")
}

// ApplyWithdrawRequest 申请提现请求
type ApplyWithdrawRequest struct {
	SubMchID   string `json:"sub_mchid"`    // 子商户号
	Amount     int    `json:"amount"`       // 提现金额（分）
	OutOrderNo string `json:"out_order_no"` // 商户订单号
}

// ApplyWithdrawResponse 申请提现响应
type ApplyWithdrawResponse struct {
	WithdrawID string `json:"withdraw_id"` // 提现单号
}
