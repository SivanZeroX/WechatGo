package api

import "github.com/wechatpy/wechatgo"

// MenuAPI 菜单管理 API
type MenuAPI struct {
	*BaseAPI
}

// NewMenuAPI 创建菜单管理 API
func NewMenuAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MenuAPI {
	return &MenuAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Get 查询自定义菜单
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html
func (api *MenuAPI) Get() (map[string]interface{}, error) {
	result, err := api.BaseAPI.Get("/menu/get", nil)
	if err != nil {
		// 菜单不存在时返回 nil
		if clientErr, ok := err.(*wechatgo.ClientError); ok {
			if clientErr.ErrCode == int(wechatgo.MenuNoExist) {
				return nil, nil
			}
		}
		return nil, err
	}
	return result, nil
}

// Create 创建自定义菜单
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
func (api *MenuAPI) Create(menuData map[string]interface{}) (map[string]interface{}, error) {
	return api.Post("/menu/create", menuData)
}

// Delete 删除自定义菜单
// https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html
func (api *MenuAPI) Delete() (map[string]interface{}, error) {
	return api.BaseAPI.Get("/menu/delete", nil)
}
