package api

import (
	"encoding/xml"
	"fmt"
)

// ProfitShareAPI 分账接口
type ProfitShareAPI struct {
	*BaseAPI
}

// NewProfitShareAPI 创建分账接口
func NewProfitShareAPI(client Client) *ProfitShareAPI {
	return &ProfitShareAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// AddProfitShare 添加分账接收方
func (api *ProfitShareAPI) AddProfitShare(req *AddProfitShareRequest) (*AddProfitShareResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":         api.client.GetAppID(),
		"mch_id":        api.client.GetMchID(),
		"nonce_str":     RandomString(32),
		"type":          req.Type,
		"account":       req.Account,
		"relation_type": req.RelationType,
	}

	if req.Name != "" {
		params["name"] = req.Name
	}

	if req.CustomRelation != "" {
		params["custom_relation"] = req.CustomRelation
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
		APIBaseURL+"pay/profitsharing/receiver/add",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add profit share: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result AddProfitShareResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// DoProfitShare 执行分账
func (api *ProfitShareAPI) DoProfitShare(req *DoProfitShareRequest) (*DoProfitShareResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":          api.client.GetAppID(),
		"mch_id":         api.client.GetMchID(),
		"nonce_str":      RandomString(32),
		"out_order_no":   req.OutOrderNo,
		"transaction_id": req.TransactionID,
		"receivers":      req.Receivers,
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
		APIBaseURL+"pay/profitsharing/order",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to do profit share: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result DoProfitShareResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryProfitShare 查询分账结果
func (api *ProfitShareAPI) QueryProfitShare(req *QueryProfitShareRequest) (*QueryProfitShareResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":          api.client.GetAppID(),
		"mch_id":         api.client.GetMchID(),
		"nonce_str":      RandomString(32),
		"out_order_no":   req.OutOrderNo,
		"transaction_id": req.TransactionID,
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
		APIBaseURL+"pay/profitsharing/orderquery",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query profit share: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryProfitShareResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// AddProfitShareRequest 添加分账接收方请求
type AddProfitShareRequest struct {
	Type           string `json:"type"`            // 分账接收方类型
	Account        string `json:"account"`         // 分账接收方账号
	Name           string `json:"name"`            // 分账接收方姓名（可选）
	RelationType   string `json:"relation_type"`   // 与分账方的关系类型
	CustomRelation string `json:"custom_relation"` // 自定义关系类型（可选）
}

// AddProfitShareResponse 添加分账接收方响应
type AddProfitShareResponse struct {
	BaseResponse
}

// DoProfitShareRequest 执行分账请求
type DoProfitShareRequest struct {
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	TransactionID string `json:"transaction_id"` // 微信订单号
	Receivers     string `json:"receivers"`      // 分账接收方列表
}

// DoProfitShareResponse 执行分账响应
type DoProfitShareResponse struct {
	BaseResponse
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	TransactionID string `json:"transaction_id"` // 微信订单号
	Status        string `json:"status"`         // 分账单状态
	Amount        int    `json:"amount"`         // 分账金额（分）
	FinishTime    string `json:"finish_time"`    // 分账完成时间
}

// QueryProfitShareRequest 查询分账结果请求
type QueryProfitShareRequest struct {
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	TransactionID string `json:"transaction_id"` // 微信订单号
}

// QueryProfitShareResponse 查询分账结果响应
type QueryProfitShareResponse struct {
	BaseResponse
	OutOrderNo    string `json:"out_order_no"`   // 商户分账单号
	TransactionID string `json:"transaction_id"` // 微信订单号
	Status        string `json:"status"`         // 分账单状态
	Amount        int    `json:"amount"`         // 分账金额（分）
	FinishTime    string `json:"finish_time"`    // 分账完成时间
	Receivers     string `json:"receivers"`      // 分账接收方列表
}
