package api

// UserAPI 用户管理 API
type UserAPI struct {
	*BaseAPI
}

// NewUserAPI 创建用户管理 API
func NewUserAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *UserAPI {
	return &UserAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Get 获取用户基本信息
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html
func (api *UserAPI) Get(openID string, lang string) (map[string]interface{}, error) {
	if lang == "" {
		lang = "zh_CN"
	}
	return api.BaseAPI.Get("/user/info", map[string]string{
		"openid": openID,
		"lang":   lang,
	})
}

// GetFollowers 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (api *UserAPI) GetFollowers(nextOpenID string) (map[string]interface{}, error) {
	params := make(map[string]string)
	if nextOpenID != "" {
		params["next_openid"] = nextOpenID
	}
	return api.BaseAPI.Get("/user/get", params)
}

// UpdateRemark 设置用户备注名
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html
func (api *UserAPI) UpdateRemark(openID, remark string) (map[string]interface{}, error) {
	return api.Post("/user/info/updateremark", map[string]interface{}{
		"openid": openID,
		"remark": remark,
	})
}

// GetBatch 批量获取用户基本信息
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html
func (api *UserAPI) GetBatch(openIDs []string, lang string) (map[string]interface{}, error) {
	if lang == "" {
		lang = "zh_CN"
	}

	userList := make([]map[string]string, len(openIDs))
	for i, openID := range openIDs {
		userList[i] = map[string]string{
			"openid": openID,
			"lang":   lang,
		}
	}

	return api.Post("/user/info/batchget", map[string]interface{}{
		"user_list": userList,
	})
}
