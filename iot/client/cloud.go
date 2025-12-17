package client

import (
	"github.com/wechatpy/wechatgo/client/api"
)

// CloudAPI 云端API
type CloudAPI struct {
	*api.BaseAPI
}

// NewCloudAPI 创建云端API
func NewCloudAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *CloudAPI {
	return &CloudAPI{
		BaseAPI: api.NewBaseAPI(client),
	}
}

// GetDeviceList 获取设备列表
// https://iot.weixin.qq.com/doc/iotdevice/cloud/queryDevice
func (api *CloudAPI) GetDeviceList(req *GetDeviceListRequest) (*GetDeviceListResponse, error) {
	result, err := api.Post("/cloud/queryDevice", req)
	if err != nil {
		return nil, err
	}

	var resp GetDeviceListResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetDeviceStatus 获取设备状态
// https://iot.weixin.qq.com/doc/iotdevice/cloud/getDeviceStatus
func (api *CloudAPI) GetDeviceStatus(req *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	result, err := api.Post("/cloud/getDeviceStatus", req)
	if err != nil {
		return nil, err
	}

	var resp GetDeviceStatusResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetDeviceControlLog 获取设备控制日志
// https://iot.weixin.qq.com/doc/iotdevice/cloud/getDeviceControlLog
func (api *CloudAPI) GetDeviceControlLog(req *GetDeviceControlLogRequest) (*GetDeviceControlLogResponse, error) {
	result, err := api.Post("/cloud/getDeviceControlLog", req)
	if err != nil {
		return nil, err
	}

	var resp GetDeviceControlLogResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ControlDevice 控制设备
// https://iot.weixin.qq.com/doc/iotdevice/cloud/controlDevice
func (api *CloudAPI) ControlDevice(req *ControlDeviceRequest) (*ControlDeviceResponse, error) {
	result, err := api.Post("/cloud/controlDevice", req)
	if err != nil {
		return nil, err
	}

	var resp ControlDeviceResponse
	if err := parseResponse(result, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
