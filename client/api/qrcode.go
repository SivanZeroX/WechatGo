package api

import (
	"net/url"

	"github.com/wechatpy/wechatgo"
)

// QRCodeAPI 二维码管理 API
type QRCodeAPI struct {
	*BaseAPI
}

// NewQRCodeAPI 创建二维码 API
func NewQRCodeAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *QRCodeAPI {
	return &QRCodeAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Create 创建二维码
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (api *QRCodeAPI) Create(qrcodeData map[string]interface{}) (map[string]interface{}, error) {
	return api.Post("/qrcode/create", qrcodeData)
}

// Scene 二维码场景信息
type Scene struct {
	SceneID   int    `json:"scene_id,omitempty"`
	SceneStr  string `json:"scene_str,omitempty"`
	SceneArgs string `json:"scene_args,omitempty"`
}

// ActionInfo 二维码动作信息
type ActionInfo struct {
	Scene *Scene `json:"scene"`
}

// QRCodeData 创建二维码的数据结构
type QRCodeData struct {
	ExpireSeconds int        `json:"expire_seconds,omitempty"`
	ActionName    string     `json:"action_name"`
	ActionInfo    ActionInfo `json:"action_info"`
}

// CreateTemporary 创建临时二维码（有效期30天）
func (api *QRCodeAPI) CreateTemporary(sceneID int, expireSeconds int) (map[string]interface{}, error) {
	data := QRCodeData{
		ExpireSeconds: expireSeconds,
		ActionName:    "QR_SCENE",
		ActionInfo: ActionInfo{
			Scene: &Scene{SceneID: sceneID},
		},
	}
	return api.Create(map[string]interface{}{
		"expire_seconds": data.ExpireSeconds,
		"action_name":    data.ActionName,
		"action_info":    data.ActionInfo,
	})
}

// CreatePermanent 创建永久二维码（数字）
func (api *QRCodeAPI) CreatePermanent(sceneID int) (map[string]interface{}, error) {
	data := QRCodeData{
		ActionName: "QR_LIMIT_SCENE",
		ActionInfo: ActionInfo{
			Scene: &Scene{SceneID: sceneID},
		},
	}
	return api.Create(map[string]interface{}{
		"action_name": data.ActionName,
		"action_info": data.ActionInfo,
	})
}

// CreatePermanentStr 创建永久二维码（字符串）
func (api *QRCodeAPI) CreatePermanentStr(sceneStr string) (map[string]interface{}, error) {
	data := QRCodeData{
		ActionName: "QR_LIMIT_STR_SCENE",
		ActionInfo: ActionInfo{
			Scene: &Scene{SceneStr: sceneStr},
		},
	}
	return api.Create(map[string]interface{}{
		"action_name": data.ActionName,
		"action_info": data.ActionInfo,
	})
}

// GetURL 通过ticket换取二维码地址
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (api *QRCodeAPI) GetURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket)
}

// Show 通过ticket获取二维码（返回图片数据）
func (api *QRCodeAPI) Show(ticket string) ([]byte, error) {
	if len(ticket) > 0 && ticket[0:1] == "{" {
		// 如果ticket是字符串形式的JSON，取其中的ticket字段
		// 这里简化处理，假设直接传入的就是ticket字符串
	}

	urlStr := api.GetURL(ticket)
	resp, err := api.httpClient.GetRaw(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	data := make([]byte, 0)
	buffer := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			data = append(data, buffer[:n]...)
		}
		if err != nil {
			break
		}
	}

	return data, nil
}

// ShowByMap 通过ticket获取二维码（从map中提取ticket）
func (api *QRCodeAPI) ShowByMap(result map[string]interface{}) ([]byte, error) {
	if ticket, ok := result["ticket"].(string); ok {
		return api.Show(ticket)
	}
	return nil, wechatgo.NewError(-1, "invalid ticket in response")
}

// GetURLByMap 通过ticket获取二维码地址（从map中提取ticket）
func (api *QRCodeAPI) GetURLByMap(result map[string]interface{}) (string, error) {
	if ticket, ok := result["ticket"].(string); ok {
		return api.GetURL(ticket), nil
	}
	return "", wechatgo.NewError(-1, "invalid ticket in response")
}
