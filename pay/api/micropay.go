package api

import (
	"encoding/xml"
	"fmt"
)

// MicroPayAPI 刷卡支付接口
type MicroPayAPI struct {
	*BaseAPI
}

// NewMicroPayAPI 创建刷卡支付接口
func NewMicroPayAPI(client Client) *MicroPayAPI {
	return &MicroPayAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Pay 刷卡支付
func (api *MicroPayAPI) Pay(req *MicroPayRequest) (*MicroPayResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":            api.client.GetAppID(),
		"mch_id":           api.client.GetMchID(),
		"nonce_str":        RandomString(32),
		"body":             req.Body,
		"out_trade_no":     req.OutTradeNo,
		"total_fee":        fmt.Sprintf("%d", req.TotalFee),
		"spbill_create_ip": req.SpbillCreateIP,
		"auth_code":        req.AuthCode,
	}

	if req.DeviceInfo != "" {
		params["device_info"] = req.DeviceInfo
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
		APIBaseURL+"pay/micropay",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to micropay: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result MicroPayResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// Reverse 撤销订单
func (api *MicroPayAPI) Reverse(req *ReverseRequest) (*ReverseResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":        api.client.GetAppID(),
		"mch_id":       api.client.GetMchID(),
		"nonce_str":    RandomString(32),
		"out_trade_no": req.OutTradeNo,
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
		APIBaseURL+"secapi/pay/reverse",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to reverse: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result ReverseResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// MicroPayRequest 刷卡支付请求
type MicroPayRequest struct {
	DeviceInfo     string `json:"device_info"`      // 设备号
	Body           string `json:"body"`             // 商品描述
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TotalFee       int    `json:"total_fee"`        // 总金额（分）
	SpbillCreateIP string `json:"spbill_create_ip"` // 终端IP
	AuthCode       string `json:"auth_code"`        // 授权码
}

// MicroPayResponse 刷卡支付响应
type MicroPayResponse struct {
	BaseResponse
	DeviceInfo  string `json:"device_info"`  // 设备号
	OpenID      string `json:"openid"`       // 用户标识
	IsSubscribe string `json:"is_subscribe"` // 是否关注公众账号
	TradeType   string `json:"trade_type"`   // 交易类型
	BankType    string `json:"bank_type"`    // 银行类型
	TotalFee    int    `json:"total_fee"`    // 总金额（分）
	CashFee     int    `json:"cash_fee"`     // 现金支付金额
	TimeEnd     string `json:"time_end"`     // 支付完成时间
}

// ReverseRequest 撤销订单请求
type ReverseRequest struct {
	OutTradeNo    string `json:"out_trade_no"`   // 商户订单号
	TransactionID string `json:"transaction_id"` // 微信订单号
}

// ReverseResponse 撤销订单响应
type ReverseResponse struct {
	BaseResponse
}
