package api

import (
	"fmt"
)

// TagAPI 标签管理 API
type TagAPI struct {
	*BaseAPI
}

// NewTagAPI 创建标签管理 API
func NewTagAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *TagAPI {
	return &TagAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Tag 标签信息
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create 创建标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) Create(name string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tag": map[string]string{
			"name": name,
		},
	}
	return api.Post("/tags/create", data)
}

// Get 获取公众号已创建的标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) Get() ([]Tag, error) {
	result, err := api.BaseAPI.Get("/tags/get", nil)
	if err != nil {
		return nil, err
	}

	if tags, ok := result["tags"].([]interface{}); ok {
		tagList := make([]Tag, len(tags))
		for i, tag := range tags {
			if tagMap, ok := tag.(map[string]interface{}); ok {
				tagList[i] = Tag{
					ID:   int(tagMap["id"].(float64)),
					Name: tagMap["name"].(string),
				}
			}
		}
		return tagList, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// Update 编辑标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) Update(tagID int, name string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tag": map[string]interface{}{
			"id":   tagID,
			"name": name,
		},
	}
	return api.Post("/tags/update", data)
}

// Delete 删除标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) Delete(tagID int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tag": map[string]interface{}{
			"id": tagID,
		},
	}
	return api.Post("/tags/delete", data)
}

// TagUser 批量为用户打标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) TagUser(tagID int, userIDs []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tagid":       tagID,
		"openid_list": userIDs,
	}
	return api.Post("/tags/members/batchtagging", data)
}

// UntagUser 批量为用户取消标签
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) UntagUser(tagID int, userIDs []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tagid":       tagID,
		"openid_list": userIDs,
	}
	return api.Post("/tags/members/batchuntagging", data)
}

// GetUserTag 获取用户身上的标签列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) GetUserTag(userID string) ([]int, error) {
	result, err := api.Post("/tags/getidlist", map[string]string{"openid": userID})
	if err != nil {
		return nil, err
	}

	if tagidList, ok := result["tagid_list"].([]interface{}); ok {
		tagIDs := make([]int, len(tagidList))
		for i, tagID := range tagidList {
			tagIDs[i] = int(tagID.(float64))
		}
		return tagIDs, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetTagUsers 获取标签下粉丝列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html
func (api *TagAPI) GetTagUsers(tagID int, firstUserID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"tagid": tagID,
	}
	if firstUserID != "" {
		data["next_openid"] = firstUserID
	}
	return api.Post("/user/tag/get", data)
}

// GetBlackList 获取公众号的黑名单列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (api *TagAPI) GetBlackList(beginOpenID string) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	if beginOpenID != "" {
		data["begin_openid"] = beginOpenID
	}
	return api.Post("/tags/members/getblacklist", data)
}

// BatchBlackList 批量拉黑用户
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (api *TagAPI) BatchBlackList(openIDList []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"openid_list": openIDList,
	}
	return api.Post("/tags/members/batchblacklist", data)
}

// BatchUnblackList 批量取消拉黑
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
func (api *TagAPI) BatchUnblackList(openIDList []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"openid_list": openIDList,
	}
	return api.Post("/tags/members/batchunblacklist", data)
}
