package api

import (
	"encoding/xml"
	"fmt"
)

// TransferAPI 企业付款接口
type TransferAPI struct {
	*BaseAPI
}

// NewTransferAPI 创建企业付款接口
func NewTransferAPI(client Client) *TransferAPI {
	return &TransferAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Transfer 企业付款
func (api *TransferAPI) Transfer(req *TransferRequest) (*TransferResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"mchid":            api.client.GetMchID(),
		"out_biz_no":       req.OutBizNo,
		"pay_scene":        req.PayScene,
		"payer_type":       "OPENID",
		"openid":           req.OpenID,
		"check_name":       req.CheckName,
		"amount":           fmt.Sprintf("%d", req.Amount),
		"appid":            api.client.GetAppID(),
		"description":      req.Description,
		"spbill_create_ip": req.SpbillCreateIP,
	}

	if req.PayerName != "" {
		params["payer_real_name"] = req.PayerName
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
		APIBaseURL+"mmpaymkttransfers/promotion/transfers",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result TransferResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryTransfer 查询企业付款
func (api *TransferAPI) QueryTransfer(req *QueryTransferRequest) (*QueryTransferResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":      api.client.GetAppID(),
		"mchid":      api.client.GetMchID(),
		"out_biz_no": req.OutBizNo,
		"nonce_str":  RandomString(32),
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
		APIBaseURL+"mmpaymkttransfers/gettransferinfo",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query transfer: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryTransferResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// TransferRequest 企业付款请求
type TransferRequest struct {
	OutBizNo       string `json:"out_biz_no"`       // 商户订单号
	PayScene       string `json:"pay_scene"`        // 支付场景
	OpenID         string `json:"openid"`           // 接收者OpenID
	CheckName      string `json:"check_name"`       // 校验方式
	PayerName      string `json:"payer_name"`       // 收款人真实姓名（可选）
	Amount         int    `json:"amount"`           // 金额（分）
	Description    string `json:"description"`      // 付款描述
	SpbillCreateIP string `json:"spbill_create_ip"` // 终端IP
}

// TransferResponse 企业付款响应
type TransferResponse struct {
	BaseResponse
	OutBizNo    string `json:"out_biz_no"`   // 商户订单号
	TransferFee int    `json:"transfer_fee"` // 手续费（分）
	PayTime     string `json:"pay_time"`     // 付款时间
}

// QueryTransferRequest 查询企业付款请求
type QueryTransferRequest struct {
	OutBizNo string `json:"out_biz_no"` // 商户订单号
}

// QueryTransferResponse 查询企业付款响应
type QueryTransferResponse struct {
	BaseResponse
	OutBizNo      string `json:"out_biz_no"`     // 商户订单号
	PayScene      string `json:"pay_scene"`      // 支付场景
	OpenID        string `json:"openid"`         // 接收者OpenID
	TransferState string `json:"transfer_state"` // 转账状态
	TransferFee   int    `json:"transfer_fee"`   // 手续费（分）
	PayTime       string `json:"pay_time"`       // 付款时间
}
