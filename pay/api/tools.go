package api

import (
	"encoding/xml"
	"fmt"
)

// ToolsAPI 工具类接口
type ToolsAPI struct {
	*BaseAPI
}

// NewToolsAPI 创建工具类接口
func NewToolsAPI(client Client) *ToolsAPI {
	return &ToolsAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// GetPublicKey 获取平台证书
func (api *ToolsAPI) GetPublicKey() (string, error) {
	// 构建请求参数
	params := map[string]string{
		"mch_id":    api.client.GetMchID(),
		"nonce_str": RandomString(32),
	}

	// 生成签名
	sign := GenerateSignature(params, api.client.GetAPIKey())
	params["sign"] = sign

	// 转换为XML
	xmlData, err := xml.Marshal(map[string]string(params))
	if err != nil {
		return "", fmt.Errorf("failed to marshal xml: %w", err)
	}

	// 发送请求
	resp, err := api.client.GetHTTPClient().Post(
		APIBaseURL+"risk/getpublickey",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to get public key: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result GetPublicKeyResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if result.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("get public key failed: %s", result.ErrCodeDes)
	}

	return result.PublicKey, nil
}

// GetPublicKeyResponse 获取平台证书响应
type GetPublicKeyResponse struct {
	BaseResponse
	ResultCode string `json:"result_code"`  // 业务结果码
	ErrCode    string `json:"err_code"`     // 错误代码
	ErrCodeDes string `json:"err_code_des"` // 错误代码描述
	PublicKey  string `json:"pub_key"`      // 平台证书
}
