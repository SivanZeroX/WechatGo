package api

import (
	"encoding/xml"
	"fmt"
)

// OrderAPI 订单接口
type OrderAPI struct {
	*BaseAPI
}

// NewOrderAPI 创建订单API
func NewOrderAPI(client Client) *OrderAPI {
	return &OrderAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// GetPrepayID 获取预支付交易会话标识
func (api *OrderAPI) GetPrepayID(req *PrepayRequest) (*PrepayResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":            req.AppID,
		"mch_id":           req.MchID,
		"body":             req.Body,
		"out_trade_no":     req.OutTradeNo,
		"total_fee":        fmt.Sprintf("%d", req.TotalFee),
		"spbill_create_ip": req.SpbillCreateIP,
		"notify_url":       req.NotifyURL,
		"trade_type":       req.TradeType,
	}

	if req.OpenID != "" {
		params["openid"] = req.OpenID
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
		APIBaseURL+"pay/unifiedorder",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request prepay: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result PrepayResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryOrder 查询订单
func (api *OrderAPI) QueryOrder(req *QueryOrderRequest) (*QueryOrderResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":  api.client.GetAppID(),
		"mch_id": api.client.GetMchID(),
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
		APIBaseURL+"pay/orderquery",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query order: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryOrderResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// CloseOrder 关闭订单
func (api *OrderAPI) CloseOrder(req *CloseOrderRequest) (*CloseOrderResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":        api.client.GetAppID(),
		"mch_id":       api.client.GetMchID(),
		"out_trade_no": req.OutTradeNo,
		"nonce_str":    RandomString(32),
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
		APIBaseURL+"pay/closeorder",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to close order: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result CloseOrderResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryOrderRequest 查询订单请求
type QueryOrderRequest struct {
	OutTradeNo    string `json:"out_trade_no"`   // 商户订单号
	TransactionID string `json:"transaction_id"` // 微信订单号
}

// QueryOrderResponse 查询订单响应
type QueryOrderResponse struct {
	BaseResponse
	DeviceInfo     string `json:"device_info"`      // 设备号
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TransactionID  string `json:"transaction_id"`   // 微信订单号
	TradeType      string `json:"trade_type"`       // 交易类型
	TradeState     string `json:"trade_state"`      // 交易状态
	TradeStateDesc string `json:"trade_state_desc"` // 交易状态描述
	BankType       string `json:"bank_type"`        // 银行类型
	TotalFee       int    `json:"total_fee"`        // 总金额（分）
	CashFee        int    `json:"cash_fee"`         // 现金支付金额
	TimeEnd        string `json:"time_end"`         // 支付完成时间
}

// CloseOrderRequest 关闭订单请求
type CloseOrderRequest struct {
	OutTradeNo string `json:"out_trade_no"` // 商户订单号
}

// CloseOrderResponse 关闭订单响应
type CloseOrderResponse struct {
	BaseResponse
}
