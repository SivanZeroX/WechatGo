package api

// TemplateAPI 模板消息和订阅通知 API
type TemplateAPI struct {
	*BaseAPI
}

// NewTemplateAPI 创建模板消息 API
func NewTemplateAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *TemplateAPI {
	return &TemplateAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// SetIndustry 设置所属行业
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#0
func (api *TemplateAPI) SetIndustry(industryID1, industryID2 string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"industry_id1": industryID1,
		"industry_id2": industryID2,
	}
	return api.Post("/template/api_set_industry", data)
}

// GetIndustry 获取设置的行业信息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#1
func (api *TemplateAPI) GetIndustry() (map[string]interface{}, error) {
	return api.BaseAPI.Get("/template/get_industry", nil)
}

// Get 获得模板ID
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#2
func (api *TemplateAPI) Get(templateIDShort string) (string, error) {
	result, err := api.Post("/template/api_add_template", map[string]interface{}{
		"template_id_short": templateIDShort,
	})
	if err != nil {
		return "", err
	}

	if templateID, ok := result["template_id"].(string); ok {
		return templateID, nil
	}
	return "", err
}

// Add Alias for Get
func (api *TemplateAPI) Add(templateIDShort string) (string, error) {
	return api.Get(templateIDShort)
}

// GetAllPrivateTemplate 获取模板列表
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#3
func (api *TemplateAPI) GetAllPrivateTemplate() (map[string]interface{}, error) {
	return api.BaseAPI.Get("/template/get_all_private_template", nil)
}

// DelPrivateTemplate 删除模板
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#4
func (api *TemplateAPI) DelPrivateTemplate(templateID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"template_id": templateID,
	}
	return api.Post("/template/del_private_template", data)
}

// AddSubscribeMessageTemplate 选用订阅通知模板
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) AddSubscribeMessageTemplate(tid string, keywords []int, description string) (string, error) {
	data := map[string]interface{}{
		"tid":       tid,
		"kidList":   keywords,
		"sceneDesc": description,
	}
	result, err := api.Post("/wxaapi/newtmpl/addtemplate", data)
	if err != nil {
		return "", err
	}

	if priTmplID, ok := result["priTmplId"].(string); ok {
		return priTmplID, nil
	}
	return "", err
}

// DelSubscribeMessageTemplate 删除订阅通知模板
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) DelSubscribeMessageTemplate(templateID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"priTmplId": templateID,
	}
	return api.Post("/wxaapi/newtmpl/deltemplate", data)
}

// SubscribeCategory 订阅通知类目
type SubscribeCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetCategory 获取公众号类目
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) GetCategory() ([]SubscribeCategory, error) {
	result, err := api.BaseAPI.Get("/wxaapi/newtmpl/getcategory", nil)
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].([]interface{}); ok {
		categories := make([]SubscribeCategory, len(data))
		for i, item := range data {
			if categoryMap, ok := item.(map[string]interface{}); ok {
				categories[i] = SubscribeCategory{
					ID:   int(categoryMap["id"].(float64)),
					Name: categoryMap["name"].(string),
				}
			}
		}
		return categories, nil
	}
	return nil, err
}

// Keyword 订阅通知关键词
type Keyword struct {
	Kid     int    `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

// GetSubscribeMessageTemplateKeywords 获取模板中的关键词
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) GetSubscribeMessageTemplateKeywords(tid string) (int, []Keyword, error) {
	result, err := api.BaseAPI.Get("/wxaapi/newtmpl/getpubtemplatekeywords", map[string]string{"tid": tid})
	if err != nil {
		return 0, nil, err
	}

	count := 0
	if countVal, ok := result["count"].(float64); ok {
		count = int(countVal)
	}

	keywords := []Keyword{}
	if data, ok := result["data"].([]interface{}); ok {
		keywords = make([]Keyword, len(data))
		for i, item := range data {
			if keywordMap, ok := item.(map[string]interface{}); ok {
				keywords[i] = Keyword{
					Kid:     int(keywordMap["kid"].(float64)),
					Name:    keywordMap["name"].(string),
					Example: keywordMap["example"].(string),
					Rule:    keywordMap["rule"].(string),
				}
			}
		}
	}

	return count, keywords, nil
}

// TemplateTitle 订阅通知模板标题
type TemplateTitle struct {
	TID        int    `json:"tid"`
	Title      string `json:"title"`
	Type       int    `json:"type"`
	CategoryID string `json:"categoryId"`
}

// GetSubscribeMessageTemplateTitles 获取所属类目的公共模板
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) GetSubscribeMessageTemplateTitles(start, limit int) (int, []TemplateTitle, error) {
	result, err := api.BaseAPI.Get("/wxaapi/newtmpl/getpubtemplatetitles", map[string]string{
		"start": string(rune(start)),
		"limit": string(rune(limit)),
	})
	if err != nil {
		return 0, nil, err
	}

	count := 0
	if countVal, ok := result["count"].(float64); ok {
		count = int(countVal)
	}

	titles := []TemplateTitle{}
	if data, ok := result["data"].([]interface{}); ok {
		titles = make([]TemplateTitle, len(data))
		for i, item := range data {
			if titleMap, ok := item.(map[string]interface{}); ok {
				titles[i] = TemplateTitle{
					TID:        int(titleMap["tid"].(float64)),
					Title:      titleMap["title"].(string),
					Type:       int(titleMap["type"].(float64)),
					CategoryID: titleMap["categoryId"].(string),
				}
			}
		}
	}

	return count, titles, nil
}

// SubscribeMessageTemplate 订阅通知私有模板
type SubscribeMessageTemplate struct {
	PriTmplID string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int    `json:"type"`
}

// GetSubscribeMessageTemplates 获取私有模板列表
// https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#addTemplate选用模板
func (api *TemplateAPI) GetSubscribeMessageTemplates() ([]SubscribeMessageTemplate, error) {
	result, err := api.BaseAPI.Get("/wxaapi/newtmpl/gettemplate", nil)
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].([]interface{}); ok {
		templates := make([]SubscribeMessageTemplate, len(data))
		for i, item := range data {
			if templateMap, ok := item.(map[string]interface{}); ok {
				templates[i] = SubscribeMessageTemplate{
					PriTmplID: templateMap["priTmplId"].(string),
					Title:     templateMap["title"].(string),
					Content:   templateMap["content"].(string),
					Example:   templateMap["example"].(string),
					Type:      int(templateMap["type"].(float64)),
				}
			}
		}
		return templates, nil
	}
	return nil, err
}
