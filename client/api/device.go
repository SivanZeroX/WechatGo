package api

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

// DeviceAPI 设备管理 API
type DeviceAPI struct {
	*BaseAPI
}

// NewDeviceAPI 创建设备管理 API
func NewDeviceAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *DeviceAPI {
	return &DeviceAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// SendMessage 主动发送消息给设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-3
func (api *DeviceAPI) SendMessage(deviceType, deviceID, userID, content string) (map[string]interface{}, error) {
	// BASE64 编码 content
	content = base64.StdEncoding.EncodeToString([]byte(content))
	data := map[string]interface{}{
		"device_type": deviceType,
		"device_id":   deviceID,
		"open_id":     userID,
		"content":     content,
	}
	return api.Post("/device/transmsg", data)
}

// SendStatusMessage 第三方主动发送设备状态消息给微信终端
// https://iot.weixin.qq.com/wiki/document-2_10.html
func (api *DeviceAPI) SendStatusMessage(deviceType, deviceID, userID string, msgType, deviceStatus int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_type":   deviceType,
		"device_id":     deviceID,
		"open_id":       userID,
		"msg_type":      msgType,
		"device_status": deviceStatus,
	}
	return api.Post("/device/transmsg", data)
}

// CreateQRCode 获取设备二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-4
func (api *DeviceAPI) CreateQRCode(deviceIDs []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_num":     len(deviceIDs),
		"device_id_list": deviceIDs,
	}
	return api.Post("/device/create_qrcode", data)
}

// GetQRCodeURL 通过 ticket 换取二维码地址
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-4
func (api *DeviceAPI) GetQRCodeURL(ticket string, data ...map[string]string) string {
	urlStr := fmt.Sprintf("https://we.qq.com/d/%s", ticket)

	if len(data) > 0 && data[0] != nil {
		// 将 map 转换为查询字符串
		params := url.Values{}
		for k, v := range data[0] {
			params.Add(k, v)
		}
		encodedData := base64.StdEncoding.EncodeToString([]byte(params.Encode()))
		urlStr = fmt.Sprintf("%s#%s", urlStr, encodedData)
	}

	return urlStr
}

// Bind 绑定设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (api *DeviceAPI) Bind(ticket, deviceID, userID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    userID,
	}
	return api.Post("/device/bind", data)
}

// Unbind 解绑设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (api *DeviceAPI) Unbind(ticket, deviceID, userID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ticket":    ticket,
		"device_id": deviceID,
		"openid":    userID,
	}
	return api.Post("/device/unbind", data)
}

// CompelBind 强制绑定用户和设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (api *DeviceAPI) CompelBind(deviceID, userID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_id": deviceID,
		"openid":    userID,
	}
	return api.Post("/device/compel_bind", data)
}

// ForceBind Alias for CompelBind
func (api *DeviceAPI) ForceBind(deviceID, userID string) (map[string]interface{}, error) {
	return api.CompelBind(deviceID, userID)
}

// CompelUnbind 强制解绑用户和设备
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-7
func (api *DeviceAPI) CompelUnbind(deviceID, userID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_id": deviceID,
		"openid":    userID,
	}
	return api.Post("/device/compel_unbind", data)
}

// ForceUnbind Alias for CompelUnbind
func (api *DeviceAPI) ForceUnbind(deviceID, userID string) (map[string]interface{}, error) {
	return api.CompelUnbind(deviceID, userID)
}

// GetStat 设备状态查询
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-8
func (api *DeviceAPI) GetStat(deviceID string) (map[string]interface{}, error) {
	return api.Get("/device/get_stat", map[string]string{"device_id": deviceID})
}

// VerifyQRCode 验证二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-9
func (api *DeviceAPI) VerifyQRCode(ticket string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"ticket": ticket,
	}
	return api.Post("/device/verify_qrcode", data)
}

// GetUserID 获取设备绑定 openID
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-11
func (api *DeviceAPI) GetUserID(deviceType, deviceID string) (map[string]interface{}, error) {
	return api.Get("/device/get_openid", map[string]string{
		"device_type": deviceType,
		"device_id":   deviceID,
	})
}

// GetOpenID Alias for GetUserID
func (api *DeviceAPI) GetOpenID(deviceType, deviceID string) (map[string]interface{}, error) {
	return api.GetUserID(deviceType, deviceID)
}

// GetBindedDevices 通过 openid 获取用户在当前 devicetype 下绑定的 deviceid 列表
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-12
func (api *DeviceAPI) GetBindedDevices(userID string) (map[string]interface{}, error) {
	return api.Get("/device/get_bind_device", map[string]string{"openid": userID})
}

// GetBindDevice Alias for GetBindedDevices
func (api *DeviceAPI) GetBindDevice(userID string) (map[string]interface{}, error) {
	return api.GetBindedDevices(userID)
}

// GetQRCode 获取 deviceid 和二维码
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-4
func (api *DeviceAPI) GetQRCode(productID ...int) (map[string]interface{}, error) {
	params := make(map[string]string)
	if len(productID) > 0 && productID[0] != 1 {
		params["product_id"] = fmt.Sprintf("%d", productID[0])
	}
	return api.Get("/device/getqrcode", params)
}

// Device 设备信息
type Device struct {
	ID          string `json:"device_id"`
	Name        string `json:"device_name"`
	Category    string `json:"category_id"`
	MAC         string `json:"mac"`
	LogoURL     string `json:"logo_url"`
	AuthType    int    `json:"auth_type"`
	OwnerName   string `json:"owner_name"`
	OwnerOpenID string `json:"openid"`
}

// Authorize 设备授权
// https://iot.weixin.qq.com/wiki/new/index.html?page=3-4-5
func (api *DeviceAPI) Authorize(devices []Device, opType int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_num":  len(devices),
		"device_list": devices,
		"op_type":     opType,
	}
	return api.Post("/device/authorize_device", data)
}
