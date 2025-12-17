package api

import (
	"encoding/xml"
	"fmt"
)

// RefundAPI 退款接口
type RefundAPI struct {
	*BaseAPI
}

// NewRefundAPI 创建退款API
func NewRefundAPI(client Client) *RefundAPI {
	return &RefundAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Refund 申请退款
func (api *RefundAPI) Refund(req *RefundRequest) (*RefundResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":           api.client.GetAppID(),
		"mch_id":          api.client.GetMchID(),
		"nonce_str":       RandomString(32),
		"out_refund_no":   req.OutRefundNo,
		"out_trade_no":    req.OutTradeNo,
		"total_fee":       fmt.Sprintf("%d", req.TotalFee),
		"refund_fee":      fmt.Sprintf("%d", req.RefundFee),
		"refund_fee_type": "CNY",
	}

	if req.TransactionID != "" {
		params["transaction_id"] = req.TransactionID
	}

	// 生成签名
	sign := GenerateSignature(params, api.client.GetAPIKey())
	params["sign"] = sign

	// 转换为XML
	xmlData, err := xml.Marshal(map[string]string(params))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal xml: %w", err)
	}

	// 发送请求
	resp, err := api.client.GetHTTPClient().Post(
		APIBaseURL+"secapi/pay/refund",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to refund: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result RefundResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryRefund 查询退款
func (api *RefundAPI) QueryRefund(req *QueryRefundRequest) (*QueryRefundResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":     api.client.GetAppID(),
		"mch_id":    api.client.GetMchID(),
		"nonce_str": RandomString(32),
	}

	if req.OutRefundNo != "" {
		params["out_refund_no"] = req.OutRefundNo
	}
	if req.RefundID != "" {
		params["refund_id"] = req.RefundID
	}
	if req.OutTradeNo != "" {
		params["out_trade_no"] = req.OutTradeNo
	}
	if req.TransactionID != "" {
		params["transaction_id"] = req.TransactionID
	}

	// 生成签名
	sign := GenerateSignature(params, api.client.GetAPIKey())
	params["sign"] = sign

	// 转换为XML
	xmlData, err := xml.Marshal(map[string]string(params))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal xml: %w", err)
	}

	// 发送请求
	resp, err := api.client.GetHTTPClient().Post(
		APIBaseURL+"pay/refundquery",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query refund: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryRefundResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// RefundRequest 退款请求
type RefundRequest struct {
	OutTradeNo    string `json:"out_trade_no"`   // 商户订单号
	OutRefundNo   string `json:"out_refund_no"`  // 商户退款单号
	TotalFee      int    `json:"total_fee"`      // 订单总金额（分）
	RefundFee     int    `json:"refund_fee"`     // 退款金额（分）
	TransactionID string `json:"transaction_id"` // 微信订单号
}

// RefundResponse 退款响应
type RefundResponse struct {
	BaseResponse
	DeviceInfo        string `json:"device_info"`         // 设备号
	TransactionID     string `json:"transaction_id"`      // 微信订单号
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	OutRefundNo       string `json:"out_refund_no"`       // 商户退款单号
	RefundID          string `json:"refund_id"`           // 微信退款单号
	RefundFee         int    `json:"refund_fee"`          // 退款金额（分）
	TotalFee          int    `json:"total_fee"`           // 订单总金额（分）
	CashFee           int    `json:"cash_fee"`            // 现金退款金额（分）
	RefundStatus      string `json:"refund_status"`       // 退款状态
	RefundRecvAccount string `json:"refund_recv_account"` // 退款入账账户
	RefundAccount     string `json:"refund_account"`      // 退款资金来源
	Time              string `json:"time"`                // 退款时间
}

// QueryRefundRequest 查询退款请求
type QueryRefundRequest struct {
	OutRefundNo   string `json:"out_refund_no"`  // 商户退款单号
	RefundID      string `json:"refund_id"`      // 微信退款单号
	OutTradeNo    string `json:"out_trade_no"`   // 商户订单号
	TransactionID string `json:"transaction_id"` // 微信订单号
}

// QueryRefundResponse 查询退款响应
type QueryRefundResponse struct {
	BaseResponse
	DeviceInfo        string `json:"device_info"`         // 设备号
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	TransactionID     string `json:"transaction_id"`      // 微信订单号
	OutRefundNo       string `json:"out_refund_no"`       // 商户退款单号
	RefundID          string `json:"refund_id"`           // 微信退款单号
	RefundFee         int    `json:"refund_fee"`          // 退款金额（分）
	TotalFee          int    `json:"total_fee"`           // 订单总金额（分）
	CashFee           int    `json:"cash_fee"`            // 现金退款金额（分）
	RefundStatus      string `json:"refund_status"`       // 退款状态
	RefundRecvAccount string `json:"refund_recv_account"` // 退款入账账户
	RefundAccount     string `json:"refund_account"`      // 退款资金来源
}
