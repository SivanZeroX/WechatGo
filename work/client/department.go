package client

import (
	"encoding/json"
	"fmt"
)

// DeptAPI 部门管理API
type DeptAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewDeptAPI 创建部门管理API
func NewDeptAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *DeptAPI {
	return &DeptAPI{
		BaseAPI: client,
	}
}

// Create 创建部门
// https://developer.work.weixin.qq.com/document/path/90213
func (api *DeptAPI) Create(req *CreateDeptRequest) (*CreateDeptResponse, error) {
	result, err := api.BaseAPI.Post("/department/create", req)
	if err != nil {
		return nil, err
	}

	var resp CreateDeptResponse
	if err := json.Unmarshal([]byte(result["id"].(string)), &resp.ID); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Get 获取部门列表
// https://developer.work.weixin.qq.com/document/path/90213
func (api *DeptAPI) Get() (*GetDeptResponse, error) {
	result, err := api.BaseAPI.Get("/department/list", nil)
	if err != nil {
		return nil, err
	}

	var resp GetDeptResponse
	if err := json.Unmarshal([]byte(result["department"].(string)), &resp.Department); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Update 更新部门
// https://developer.work.weixin.qq.com/document/path/90213
func (api *DeptAPI) Update(req *UpdateDeptRequest) error {
	_, err := api.BaseAPI.Post("/department/update", req)
	return err
}

// Delete 删除部门
// https://developer.work.weixin.qq.com/document/path/90213
func (api *DeptAPI) Delete(deptID int) error {
	_, err := api.BaseAPI.Get("/department/delete", map[string]string{
		"id": fmt.Sprintf("%d", deptID),
	})
	return err
}

// ==================== 请求结构体 ====================

// CreateDeptRequest 创建部门请求
type CreateDeptRequest struct {
	Name     string `json:"name"`     // 部门名称
	ParentID int    `json:"parentid"` // 父部门ID
	Order    int    `json:"order"`    // 部门排序
}

// UpdateDeptRequest 更新部门请求
type UpdateDeptRequest struct {
	ID       int    `json:"id"`       // 部门ID
	Name     string `json:"name"`     // 部门名称
	ParentID int    `json:"parentid"` // 父部门ID
	Order    int    `json:"order"`    // 部门排序
}

// ==================== 响应结构体 ====================

// CreateDeptResponse 创建部门响应
type CreateDeptResponse struct {
	ID int `json:"id"` // 部门ID
}

// GetDeptResponse 获取部门列表响应
type GetDeptResponse struct {
	Department []Dept `json:"department"` // 部门列表
}

// Dept 部门信息
type Dept struct {
	ID       int    `json:"id"`       // 部门ID
	Name     string `json:"name"`     // 部门名称
	ParentID int    `json:"parentid"` // 父部门ID
	Order    int    `json:"order"`    // 部门排序
}
