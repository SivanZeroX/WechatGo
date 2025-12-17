package client

import (
	"encoding/json"
	"fmt"
)

// TagAPI 标签管理API
type TagAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewTagAPI 创建标签管理API
func NewTagAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *TagAPI {
	return &TagAPI{
		BaseAPI: client,
	}
}

// Create 创建标签
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) Create(req *CreateTagRequest) (*CreateTagResponse, error) {
	result, err := api.BaseAPI.Post("/tag/create", req)
	if err != nil {
		return nil, err
	}

	var resp CreateTagResponse
	if err := json.Unmarshal([]byte(result["tagid"].(string)), &resp.TagID); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Get 获取标签列表
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) Get() (*GetTagResponse, error) {
	result, err := api.BaseAPI.Get("/tag/list", nil)
	if err != nil {
		return nil, err
	}

	var resp GetTagResponse
	if err := json.Unmarshal([]byte(result["taglist"].(string)), &resp.TagList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// Update 更新标签
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) Update(req *UpdateTagRequest) error {
	_, err := api.BaseAPI.Post("/tag/update", req)
	return err
}

// Delete 删除标签
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) Delete(tagID int) error {
	_, err := api.BaseAPI.Get("/tag/delete", map[string]string{
		"tagid": fmt.Sprintf("%d", tagID),
	})
	return err
}

// AddTagUsers 添加标签成员
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) AddTagUsers(tagID int, userIDs []string) error {
	req := map[string]interface{}{
		"tagid":    tagID,
		"userlist": userIDs,
	}
	_, err := api.BaseAPI.Post("/tag/addtagusers", req)
	return err
}

// DelTagUsers 删除标签成员
// https://developer.work.weixin.qq.com/document/path/90219
func (api *TagAPI) DelTagUsers(tagID int, userIDs []string) error {
	req := map[string]interface{}{
		"tagid":    tagID,
		"userlist": userIDs,
	}
	_, err := api.BaseAPI.Post("/tag/deltagusers", req)
	return err
}

// ==================== 请求结构体 ====================

// CreateTagRequest 创建标签请求
type CreateTagRequest struct {
	TagName string `json:"tagname"` // 标签名称
}

// UpdateTagRequest 更新标签请求
type UpdateTagRequest struct {
	TagID   int    `json:"tagid"`   // 标签ID
	TagName string `json:"tagname"` // 标签名称
}

// ==================== 响应结构体 ====================

// CreateTagResponse 创建标签响应
type CreateTagResponse struct {
	TagID int `json:"tagid"` // 标签ID
}

// GetTagResponse 获取标签列表响应
type GetTagResponse struct {
	TagList []Tag `json:"taglist"` // 标签列表
}

// Tag 标签信息
type Tag struct {
	TagID   int    `json:"tagid"`   // 标签ID
	TagName string `json:"tagname"` // 标签名称
}
