package client

import (
	"github.com/wechatpy/wechatgo/client/api"
)

// DeviceAPI 设备API
type DeviceAPI struct {
	*api.BaseAPI
}

// NewDeviceAPI 创建设备API
func NewDeviceAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *DeviceAPI {
	return &DeviceAPI{
		BaseAPI: api.NewBaseAPI(client),
	}
}

// ApplyDevice 申请设备
// https://iot.weixin.qq.com/doc/iotdevice/device/applyDevice
func (api *DeviceAPI) ApplyDevice(req *ApplyDeviceRequest) (*ApplyDeviceResponse, error) {
	result, err := api.Post("/device/applyDevice", req)
	if err != nil {
		return nil, err
	}

	var resp ApplyDeviceResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetDeviceQRCode 获取设备二维码
// https://iot.weixin.qq.com/doc/iotdevice/device/getDeviceQRCode
func (api *DeviceAPI) GetDeviceQRCode(req *GetDeviceQRCodeRequest) (*GetDeviceQRCodeResponse, error) {
	result, err := api.Post("/device/getDeviceQRCode", req)
	if err != nil {
		return nil, err
	}

	var resp GetDeviceQRCodeResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// BindDevice 绑定设备
// https://iot.weixin.qq.com/doc/iotdevice/device/bindDevice
func (api *DeviceAPI) BindDevice(req *BindDeviceRequest) (*BindDeviceResponse, error) {
	result, err := api.Post("/device/bindDevice", req)
	if err != nil {
		return nil, err
	}

	var resp BindDeviceResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnbindDevice 解绑设备
// https://iot.weixin.qq.com/doc/iotdevice/device/unbindDevice
func (api *DeviceAPI) UnbindDevice(req *UnbindDeviceRequest) (*UnbindDeviceResponse, error) {
	result, err := api.Post("/device/unbindDevice", req)
	if err != nil {
		return nil, err
	}

	var resp UnbindDeviceResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
