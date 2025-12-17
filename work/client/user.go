package client

import (
	"encoding/json"
	"fmt"
)

// UserAPI 用户管理API
type UserAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewUserAPI 创建用户管理API
func NewUserAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *UserAPI {
	return &UserAPI{
		BaseAPI: client,
	}
}

// Create 创建用户
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) Create(req *CreateUserRequest) error {
	_, err := api.BaseAPI.Post("/user/create", req)
	return err
}

// Get 获取用户
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) Get(userID string) (*GetUserResponse, error) {
	result, err := api.BaseAPI.Get("/user/get", map[string]string{"userid": userID})
	if err != nil {
		return nil, err
	}

	var resp GetUserResponse
	if err := json.Unmarshal([]byte(result["userlist"].(string)), &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Update 更新用户
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) Update(userID string, req *UpdateUserRequest) error {
	_, err := api.BaseAPI.Post("/user/update", req)
	return err
}

// Delete 删除用户
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) Delete(userID string) error {
	_, err := api.BaseAPI.Get("/user/delete", map[string]string{"userid": userID})
	return err
}

// BatchDelete 批量删除用户
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) BatchDelete(userIDs []string) error {
	req := map[string]interface{}{
		"useridlist": userIDs,
	}
	_, err := api.BaseAPI.Post("/user/batchdelete", req)
	return err
}

// SimpleList 获取部门成员
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) SimpleList(deptID int, fetchChild bool) (*SimpleListResponse, error) {
	result, err := api.BaseAPI.Get("/user/simplelist", map[string]string{
		"department_id": fmt.Sprintf("%d", deptID),
		"fetch_child":   fmt.Sprintf("%t", fetchChild),
	})
	if err != nil {
		return nil, err
	}

	var resp SimpleListResponse
	if err := json.Unmarshal([]byte(result["userlist"].(string)), &resp.UserList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// List 获取部门成员详情
// https://developer.work.weixin.qq.com/document/path/90196
func (api *UserAPI) List(deptID int, fetchChild bool) (*ListResponse, error) {
	result, err := api.BaseAPI.Get("/user/list", map[string]string{
		"department_id": fmt.Sprintf("%d", deptID),
		"fetch_child":   fmt.Sprintf("%t", fetchChild),
	})
	if err != nil {
		return nil, err
	}

	var resp ListResponse
	if err := json.Unmarshal([]byte(result["userlist"].(string)), &resp.UserList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// ==================== 请求结构体 ====================

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	UserID     string `json:"userid"`      // 成员UserID
	Name       string `json:"name"`        // 成员名称
	Department []int  `json:"department"`  // 部门ID列表
	Mobile     string `json:"mobile"`      // 手机号
	Email      string `json:"email"`       // 邮箱
	Position   string `json:"position"`    // 职位
	Telephone  string `json:"telephone"`   // 座机
	Alias      string `json:"alias"`       // 别名
	OpenUserID string `json:"open_userid"` // 外部用户ID
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	UserID     string `json:"userid"`     // 成员UserID
	Name       string `json:"name"`       // 成员名称
	Department []int  `json:"department"` // 部门ID列表
	Mobile     string `json:"mobile"`     // 手机号
	Email      string `json:"email"`      // 邮箱
	Position   string `json:"position"`   // 职位
	Telephone  string `json:"telephone"`  // 座机
	Alias      string `json:"alias"`      // 别名
}

// ==================== 响应结构体 ====================

// GetUserResponse 获取用户响应
type GetUserResponse struct {
	UserID     string `json:"userid"`      // 成员UserID
	Name       string `json:"name"`        // 成员名称
	Department []int  `json:"department"`  // 部门ID列表
	Mobile     string `json:"mobile"`      // 手机号
	Email      string `json:"email"`       // 邮箱
	Position   string `json:"position"`    // 职位
	Telephone  string `json:"telephone"`   // 座机
	Alias      string `json:"alias"`       // 别名
	OpenUserID string `json:"open_userid"` // 外部用户ID
	Avatar     string `json:"avatar"`      // 头像
	Status     int    `json:"status"`      // 激活状态
}

// SimpleListResponse 获取部门成员响应
type SimpleListResponse struct {
	UserList []SimpleUser `json:"userlist"` // 成员列表
}

// SimpleUser 简单用户信息
type SimpleUser struct {
	UserID     string `json:"userid"`     // 成员UserID
	Name       string `json:"name"`       // 成员名称
	Mobile     string `json:"mobile"`     // 手机号
	Email      string `json:"email"`      // 邮箱
	Department []int  `json:"department"` // 部门ID列表
}

// ListResponse 获取部门成员详情响应
type ListResponse struct {
	UserList []User `json:"userlist"` // 成员列表
}

// User 用户详细信息
type User struct {
	UserID     string `json:"userid"`      // 成员UserID
	Name       string `json:"name"`        // 成员名称
	Department []int  `json:"department"`  // 部门ID列表
	Mobile     string `json:"mobile"`      // 手机号
	Email      string `json:"email"`       // 邮箱
	Position   string `json:"position"`    // 职位
	Telephone  string `json:"telephone"`   // 座机
	Alias      string `json:"alias"`       // 别名
	OpenUserID string `json:"open_userid"` // 外部用户ID
	Avatar     string `json:"avatar"`      // 头像
	Status     int    `json:"status"`      // 激活状态
	Gender     int    `json:"gender"`      // 性别
}
