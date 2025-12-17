package api

// POIAPI 门店管理 API
type POIAPI struct {
	*BaseAPI
}

// NewPOIAPI 创建门店管理 API
func NewPOIAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *POIAPI {
	return &POIAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Add 创建门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#7
func (api *POIAPI) Add(poiData map[string]interface{}) (map[string]interface{}, error) {
	return api.Post("/poi/addpoi", poiData)
}

// Get 查询门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#9
func (api *POIAPI) Get(poiID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"poi_id": poiID,
	}
	return api.Post("/poi/getpoi", data)
}

// List 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#10
func (api *POIAPI) List(begin, limit int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"begin": begin,
		"limit": limit,
	}
	return api.Post("/poi/getpoilist", data)
}

// Update 修改门店服务信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#11
func (api *POIAPI) Update(poiData map[string]interface{}) (map[string]interface{}, error) {
	return api.Post("/poi/updatepoi", poiData)
}

// Delete 删除门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#12
func (api *POIAPI) Delete(poiID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"poi_id": poiID,
	}
	return api.Post("/poi/delpoi", data)
}

// Category 门店类目
type Category struct {
	ClassName string     `json:"class_name"`
	ClassID   int        `json:"class_id"`
	ParentID  int        `json:"parent_id"`
	Level     int        `json:"level"`
	Children  []Category `json:"category_list,omitempty"`
}

// GetCategories 获取微信门店类目表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html#13
func (api *POIAPI) GetCategories() ([]Category, error) {
	result, err := api.BaseAPI.Get("/cgi-bin/api_getwxcategory", nil)
	if err != nil {
		return nil, err
	}

	if categoryList, ok := result["category_list"].([]interface{}); ok {
		categories := make([]Category, len(categoryList))
		for i, item := range categoryList {
			if categoryMap, ok := item.(map[string]interface{}); ok {
				categories[i] = Category{
					ClassName: categoryMap["class_name"].(string),
					ClassID:   int(categoryMap["class_id"].(float64)),
					ParentID:  int(categoryMap["parent_id"].(float64)),
					Level:     int(categoryMap["level"].(float64)),
				}

				// 处理子类目
				if childList, ok := categoryMap["category_list"].([]interface{}); ok {
					children := make([]Category, len(childList))
					for j, child := range childList {
						if childMap, ok := child.(map[string]interface{}); ok {
							children[j] = Category{
								ClassName: childMap["class_name"].(string),
								ClassID:   int(childMap["class_id"].(float64)),
								ParentID:  int(childMap["parent_id"].(float64)),
								Level:     int(childMap["level"].(float64)),
							}
						}
					}
					categories[i].Children = children
				}
			}
		}
		return categories, nil
	}
	return nil, err
}
