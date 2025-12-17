package client

import "time"

// ErrorResponse 错误响应
type ErrorResponse struct {
	BaseResponse
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

// BaseResponse 基础响应
type BaseResponse struct {
	Code    int    `json:"code"`    // 返回码，0为成功
	Message string `json:"message"` // 返回信息
}

// BaseRequest 基础请求
type BaseRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
}

// ==================== 云端API请求/响应 ====================

// GetDeviceListRequest 获取设备列表请求
type GetDeviceListRequest struct {
	Limit  int `json:"limit"`  // 返回数量
	Offset int `json:"offset"` // 偏移量
}

// GetDeviceListResponse 获取设备列表响应
type GetDeviceListResponse struct {
	BaseResponse
	Data struct {
		DeviceList []DeviceInfo `json:"deviceList"` // 设备列表
		Total      int          `json:"total"`      // 总数
	} `json:"data"`
}

// DeviceInfo 设备信息
type DeviceInfo struct {
	DeviceID   string `json:"deviceId"`   // 设备ID
	DeviceType string `json:"deviceType"` // 设备类型
	NickName   string `json:"nickName"`   // 设备昵称
	BindUserID string `json:"bindUserId"` // 绑定用户ID
	BindTime   string `json:"bindTime"`   // 绑定时间
	Status     int    `json:"status"`     // 设备状态 0:离线 1:在线
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

// GetDeviceStatusRequest 获取设备状态请求
type GetDeviceStatusRequest struct {
	BaseRequest
}

// GetDeviceStatusResponse 获取设备状态响应
type GetDeviceStatusResponse struct {
	BaseResponse
	Data struct {
		DeviceID  string `json:"deviceId"`  // 设备ID
		Status    int    `json:"status"`    // 设备状态 0:离线 1:在线
		OnlineAt  string `json:"onlineAt"`  // 上线时间
		OfflineAt string `json:"offlineAt"` // 离线时间
	} `json:"data"`
}

// GetDeviceControlLogRequest 获取设备控制日志请求
type GetDeviceControlLogRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
	Start    string `json:"start"`    // 开始时间 YYYY-MM-DD HH:mm:ss
	End      string `json:"end"`      // 结束时间 YYYY-MM-DD HH:mm:ss
	Limit    int    `json:"limit"`    // 返回数量
	Offset   int    `json:"offset"`   // 偏移量
}

// GetDeviceControlLogResponse 获取设备控制日志响应
type GetDeviceControlLogResponse struct {
	BaseResponse
	Data struct {
		Logs  []DeviceControlLog `json:"logs"`  // 控制日志列表
		Total int                `json:"total"` // 总数
	} `json:"data"`
}

// DeviceControlLog 设备控制日志
type DeviceControlLog struct {
	ID        int       `json:"id"`        // 日志ID
	DeviceID  string    `json:"deviceId"`  // 设备ID
	Action    string    `json:"action"`    // 操作类型
	Params    string    `json:"params"`    // 操作参数
	Result    string    `json:"result"`    // 操作结果
	CreatedAt time.Time `json:"createdAt"` // 创建时间
}

// ControlDeviceRequest 控制设备请求
type ControlDeviceRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
	Action   string `json:"action"`   // 操作类型
	Params   string `json:"params"`   // 操作参数，JSON格式
}

// ControlDeviceResponse 控制设备响应
type ControlDeviceResponse struct {
	BaseResponse
	Data struct {
		RequestID string `json:"requestId"` // 请求ID
		Result    string `json:"result"`    // 执行结果
	} `json:"data"`
}

// ==================== 设备API请求/响应 ====================

// ApplyDeviceRequest 申请设备请求
type ApplyDeviceRequest struct {
	DeviceType string `json:"deviceType"` // 设备类型
	NickName   string `json:"nickName"`   // 设备昵称
}

// ApplyDeviceResponse 申请设备响应
type ApplyDeviceResponse struct {
	BaseResponse
	Data struct {
		DeviceID string `json:"deviceId"` // 设备ID
	} `json:"data"`
}

// GetDeviceQRCodeRequest 获取设备二维码请求
type GetDeviceQRCodeRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
}

// GetDeviceQRCodeResponse 获取设备二维码响应
type GetDeviceQRCodeResponse struct {
	BaseResponse
	Data struct {
		QRCode string `json:"qrCode"` // 二维码内容
	} `json:"data"`
}

// BindDeviceRequest 绑定设备请求
type BindDeviceRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
	UserID   string `json:"userId"`   // 用户ID
}

// BindDeviceResponse 绑定设备响应
type BindDeviceResponse struct {
	BaseResponse
	Data struct {
		BindTime string `json:"bindTime"` // 绑定时间
	} `json:"data"`
}

// UnbindDeviceRequest 解绑设备请求
type UnbindDeviceRequest struct {
	DeviceID string `json:"deviceId"` // 设备ID
	UserID   string `json:"userId"`   // 用户ID
}

// UnbindDeviceResponse 解绑设备响应
type UnbindDeviceResponse struct {
	BaseResponse
	Data struct {
		UnbindTime string `json:"unbindTime"` // 解绑时间
	} `json:"data"`
}

// ==================== 错误定义 ====================

var (
	ErrAccessTokenNotFound    = NewError(40001, "access_token not found")
	ErrAccessTokenInvalidType = NewError(40002, "access_token is not a string")
)

// Error 错误
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// NewError 创建错误
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// parseResponse 解析响应
func parseResponse(result map[string]interface{}, target interface{}) error {
	// TODO: 实现响应解析逻辑
	return nil
}
