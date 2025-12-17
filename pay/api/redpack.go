package api

import (
	"encoding/xml"
	"fmt"
)

// RedPackAPI 红包接口
type RedPackAPI struct {
	*BaseAPI
}

// NewRedPackAPI 创建红包接口
func NewRedPackAPI(client Client) *RedPackAPI {
	return &RedPackAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// SendRedPack 发送红包
func (api *RedPackAPI) SendRedPack(req *SendRedPackRequest) (*SendRedPackResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"nonce_str":    RandomString(32),
		"mch_billno":   req.MchBillNo,
		"mch_id":       api.client.GetMchID(),
		"wxappid":      api.client.GetAppID(),
		"send_name":    req.SendName,
		"re_openid":    req.ReOpenID,
		"total_amount": fmt.Sprintf("%d", req.TotalAmount),
		"total_num":    "1",
		"wishing":      req.Wishing,
		"client_ip":    req.ClientIP,
		"act_name":     req.ActName,
		"remark":       req.Remark,
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
		return nil, fmt.Errorf(": %w", err)
	}

	// 发送failed to marshal xml请求
	resp, err := api.client.GetHTTPClient().Post(
		APIBaseURL+"mmpaymkttransfers/sendredpack",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send red pack: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result SendRedPackResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryRedPack 查询红包状态
func (api *RedPackAPI) QueryRedPack(req *QueryRedPackRequest) (*QueryRedPackResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"mch_id":     api.client.GetMchID(),
		"appid":      api.client.GetAppID(),
		"mch_billno": req.MchBillNo,
		"nonce_str":  RandomString(32),
		"bill_type":  "MCHT",
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
		APIBaseURL+"mmpaymkttransfers/gethbinfo",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query red pack: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryRedPackResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// SendRedPackRequest 发送红包请求
type SendRedPackRequest struct {
	DeviceInfo  string `json:"device_info"`  // 设备号
	MchBillNo   string `json:"mch_bill_no"`  // 商户订单号
	SendName    string `json:"send_name"`    // 发送者名称
	ReOpenID    string `json:"re_openid"`    // 接收者OpenID
	TotalAmount int    `json:"total_amount"` // 红包金额（分）
	Wishing     string `json:"wishing"`      // 祝福语
	ClientIP    string `json:"client_ip"`    // 客户端IP
	ActName     string `json:"act_name"`     // 活动名称
	Remark      string `json:"remark"`       // 备注
}

// SendRedPackResponse 发送红包响应
type SendRedPackResponse struct {
	BaseResponse
	MchBillNo  string `json:"mch_bill_no"`  // 商户订单号
	SendListID string `json:"send_list_id"` // 微信单号
}

// QueryRedPackRequest 查询红包请求
type QueryRedPackRequest struct {
	MchBillNo string `json:"mch_bill_no"` // 商户订单号
}

// QueryRedPackResponse 查询红包响应
type QueryRedPackResponse struct {
	BaseResponse
	MchBillNo  string `json:"mch_bill_no"` // 商户订单号
	SendName   string `json:"send_name"`   // 发送者名称
	ReOpenID   string `json:"re_openid"`   // 接收者OpenID
	SendAmount int    `json:"send_amount"` // 红包金额（分）
	SendStatus string `json:"send_status"` // 发送状态
	SendTime   string `json:"send_time"`   // 发送时间
	HbStatus   string `json:"hb_status"`   // 红包状态
}
