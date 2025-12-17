package api

import (
	"fmt"
)

// BanksAPI 银行相关API
type BanksAPI struct {
	*BaseAPI
}

// NewBanksAPI 创建银行API
func NewBanksAPI(client *Client) *BanksAPI {
	return &BanksAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// GetBankList 获取银行列表
func (api *BanksAPI) GetBankList() (*BankListResponse, error) {
	// TODO: 实现获取银行列表逻辑
	return nil, fmt.Errorf("not implemented")
}

// Bank 银行信息
type Bank struct {
	BankID   string `json:"bank_id"`   // 银行ID
	BankName string `json:"bank_name"` // 银行名称
}

// BankListResponse 银行列表响应
type BankListResponse struct {
	BankList []Bank `json:"bank_list"` // 银行列表
}
