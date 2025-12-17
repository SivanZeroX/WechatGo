package wechatgo

import "encoding/xml"

// MessageType 消息类型
type MessageType string

const (
	MsgTypeText            MessageType = "text"
	MsgTypeImage           MessageType = "image"
	MsgTypeVoice           MessageType = "voice"
	MsgTypeVideo           MessageType = "video"
	MsgTypeShortVideo      MessageType = "shortvideo"
	MsgTypeLocation        MessageType = "location"
	MsgTypeLink            MessageType = "link"
	MsgTypeMiniProgramPage MessageType = "miniprogrampage"
	MsgTypeEvent           MessageType = "event"
)

// BaseMessage 基础消息
type BaseMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MsgID        int64    `xml:"MsgId,omitempty"`
}

// TextMessage 文本消息
type TextMessage struct {
	BaseMessage
	Content string `xml:"Content"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	BaseMessage
	PicURL  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	BaseMessage
	MediaID     string `xml:"MediaId"`
	Format      string `xml:"Format"`
	Recognition string `xml:"Recognition,omitempty"`
}

// VideoMessage 视频消息
type VideoMessage struct {
	BaseMessage
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

// ShortVideoMessage 短视频消息
type ShortVideoMessage struct {
	BaseMessage
	MediaID      string `xml:"MediaId"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

// LocationMessage 地理位置消息
type LocationMessage struct {
	BaseMessage
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     int     `xml:"Scale"`
	Label     string  `xml:"Label"`
}

// LinkMessage 链接消息
type LinkMessage struct {
	BaseMessage
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	URL         string `xml:"Url"`
}

// MiniProgramPageMessage 小程序卡片消息
type MiniProgramPageMessage struct {
	BaseMessage
	AppID        string `xml:"AppId"`
	Title        string `xml:"Title"`
	PagePath     string `xml:"PagePath"`
	ThumbURL     string `xml:"ThumbUrl"`
	ThumbMediaID string `xml:"ThumbMediaId"`
}

// Message 通用消息接口
type Message interface {
	GetMsgType() string
	GetFromUserName() string
	GetToUserName() string
	GetCreateTime() int64
}

// GetMsgType 获取消息类型
func (m *BaseMessage) GetMsgType() string {
	return m.MsgType
}

// GetFromUserName 获取发送者
func (m *BaseMessage) GetFromUserName() string {
	return m.FromUserName
}

// GetToUserName 获取接收者
func (m *BaseMessage) GetToUserName() string {
	return m.ToUserName
}

// GetCreateTime 获取创建时间
func (m *BaseMessage) GetCreateTime() int64 {
	return m.CreateTime
}
