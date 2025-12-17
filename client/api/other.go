package api

import (
	"fmt"
)

// MiscAPI 杂项 API
type MiscAPI struct {
	*BaseAPI
}

// NewMiscAPI 创建杂项 API
func NewMiscAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MiscAPI {
	return &MiscAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// ShortURL 将一条长链接转成短链接
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/URL_Shortener.html
//
// 注意：该接口即将废弃
// https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&announce_id=11615366683l3hgk&version=63010043&lang=zh_CN&token=
func (api *MiscAPI) ShortURL(longURL string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"action":   "long2short",
		"long_url": longURL,
	}
	return api.Post("/shorturl", data)
}

// GetWeChatIPs 获取微信服务器 IP 地址列表
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
func (api *MiscAPI) GetWeChatIPs() ([]string, error) {
	result, err := api.Get("/getcallbackip", nil)
	if err != nil {
		return nil, err
	}

	if ipList, ok := result["ip_list"].([]interface{}); ok {
		ips := make([]string, len(ipList))
		for i, ip := range ipList {
			ips[i] = ip.(string)
		}
		return ips, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// CheckNetwork 网络检测
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Network_Detection.html
func (api *MiscAPI) CheckNetwork(action, operator string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"action":         action,
		"check_operator": operator,
	}
	return api.Post("/callback/check", data)
}
