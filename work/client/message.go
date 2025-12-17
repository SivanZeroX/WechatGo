package client

import (
	"encoding/json"
	"fmt"
)

// MessageAPI 消息管理API
type MessageAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewMessageAPI 创建消息管理API
func NewMessageAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MessageAPI {
	return &MessageAPI{
		BaseAPI: client,
	}
}

// Send 发送消息
// https://developer.work.weixin.qq.com/document/path/90235
func (api *MessageAPI) Send(req *SendMessageRequest) (*SendMessageResponse, error) {
	result, err := api.BaseAPI.Post("/message/send", req)
	if err != nil {
		return nil, err
	}

	var resp SendMessageResponse
	if err := json.Unmarshal([]byte(result["msgid"].(string)), &resp.MsgID); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// SendToUser 发送消息给用户
func (api *MessageAPI) SendToUser(userIDs []string, msg interface{}) (*SendMessageResponse, error) {
	return api.Send(&SendMessageRequest{
		Touser:  userIDs,
		MsgType: "text",
		Text:    map[string]string{"content": msg.(string)},
	})
}

// ==================== 请求结构体 ====================

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Touser  []string          `json:"touser"`  // 接收者UserID列表
	Toparty []int             `json:"toparty"` // 接收部门ID列表
	Totag   []int             `json:"totag"`   // 接收标签ID列表
	MsgType string            `json:"msgtype"` // 消息类型
	AgentID string            `json:"agentid"` // 应用ID
	Text    map[string]string `json:"text"`    // 文本消息内容
	Media   map[string]string `json:"media"`   // 媒体消息内容
}

// ==================== 响应结构体 ====================

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	MsgID int64 `json:"msgid"` // 消息ID
}
