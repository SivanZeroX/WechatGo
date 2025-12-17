package client

import (
	"encoding/json"
	"fmt"
)

// ContactAPI 客户联系API
type ContactAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewContactAPI 创建客户联系API
func NewContactAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *ContactAPI {
	return &ContactAPI{
		BaseAPI: client,
	}
}

// Add 添加客户
// https://developer.work.weixin.qq.com/document/path/92125
func (api *ContactAPI) Add(req *AddContactRequest) (*AddContactResponse, error) {
	result, err := api.BaseAPI.Post("/crm/contact/add", req)
	if err != nil {
		return nil, err
	}

	var resp AddContactResponse
	if err := json.Unmarshal([]byte(result["open_userid"].(string)), &resp.OpenUserID); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Get 获取客户详情
// https://developer.work.weixin.qq.com/document/path/92125
func (api *ContactAPI) Get(openUserID string) (*GetContactResponse, error) {
	result, err := api.BaseAPI.Get("/crm/contact/get", map[string]string{
		"open_userid": openUserID,
	})
	if err != nil {
		return nil, err
	}

	var resp GetContactResponse
	if err := json.Unmarshal([]byte(result["contact"].(string)), &resp.Contact); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Update 更新客户信息
// https://developer.work.weixin.qq.com/document/path/92125
func (api *ContactAPI) Update(req *UpdateContactRequest) error {
	_, err := api.BaseAPI.Post("/crm/contact/update", req)
	return err
}

// ==================== 请求结构体 ====================

// AddContactRequest 添加客户请求
type AddContactRequest struct {
	ExternalUserID string `json:"external_userid"` // 外部用户ID
	Name           string `json:"name"`            // 客户名称
	Mobile         string `json:"mobile"`          // 手机号
	Email          string `json:"email"`           // 邮箱
}

// UpdateContactRequest 更新客户请求
type UpdateContactRequest struct {
	ExternalUserID string `json:"external_userid"` // 外部用户ID
	Name           string `json:"name"`            // 客户名称
	Mobile         string `json:"mobile"`          // 手机号
	Email          string `json:"email"`           // 邮箱
}

// ==================== 响应结构体 ====================

// AddContactResponse 添加客户响应
type AddContactResponse struct {
	OpenUserID string `json:"open_userid"` // 外部用户ID
}

// GetContactResponse 获取客户详情响应
type GetContactResponse struct {
	Contact Contact `json:"contact"` // 客户信息
}

// Contact 客户信息
type Contact struct {
	ExternalUserID string `json:"external_userid"` // 外部用户ID
	Name           string `json:"name"`            // 客户名称
	Mobile         string `json:"mobile"`          // 手机号
	Email          string `json:"email"`           // 邮箱
	Gender         int    `json:"gender"`          // 性别
}
