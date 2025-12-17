package wechatgo

/*
================================================================================
WECHATGO - WECHAT PYTHON TO GO CONVERSION PROJECT
================================================================================

Project Overview:
-----------------
This is the WeChatGo SDK - a complete Go language implementation of the
wechatpy Python SDK (v2.0.0.alpha26).

Source Project:  wechatpy - WeChat SDK for Python
Target Project:  wechatgo - WeChat SDK for Go
License:         MIT
Conversion Date: 2025-12-16

Directory Structure Mapping:
---------------------------
Python Source Structure → Go Target Structure

wechatpy/                    →  wechatgo/
├── messages.py              →  messages.go (Message type definitions)
├── events.py                →  events.go (Event type definitions)
├── parser.py   [THIS FILE]  →  wechat_parser.go (XML message parser)
├── replies.py               →  replies.go (Message reply builder)

Conversion Strategy - Phase 3: Message Processing
-------------------------------------------------
Goal: Implement message parsing and reply functionality

1. Message Types (messages.go, events.go)
   - Define message structures
   - Implement XML serialization/deserialization

2. Message Parsing (wechat_parser.go) [CURRENT FILE]
   - XML parsing
   - Message type recognition
   - Encrypted message decryption

3. Message Replies (replies.go)
   - Reply message building
   - XML serialization
   - Encrypted replies

Language Feature Mapping (Python → Go):
---------------------------------------
| Python Feature     | Go Implementation              |
|--------------------|--------------------------------|
| Class              | Struct + Methods               |
| Inheritance        | Composition + Interface        |
| Decorators         | Function wrapping/Middleware   |
| Properties         | Getter/Setter methods          |
| Exceptions         | Error return values            |
| Optional parameters| Variadic params or Options     |
| Dictionary         | map[string]interface{}         |
| None               | nil or zero value              |
| Dynamic types      | interface{} + type assertion   |

Error Handling Strategy:
-----------------------
Python Style:
    try:
        result = client.get_user(openid)
    except WeChatClientException as e:
        handle_error(e)

Go Style:
    result, err := client.GetUser(openid)
    if err != nil {
        handleError(err)
    }

Key Design Decisions:
--------------------
1. Package Structure
   - Each submodule is an independent package
   - Avoid circular dependencies
   - Clear exported interfaces

2. Interface Design
   - SessionStorage interface for session management
   - MessageParser interface for message parsing

3. Error Handling
   - Use errors.New() and fmt.Errorf()
   - Define custom error types
   - Error wrapping: fmt.Errorf("failed: %w", err)

4. HTTP Client
   - Use Go standard library net/http
   - BaseClient with token management
   - Automatic retry on token expiration

5. Concurrent Safety
   - Access Token cache uses sync.RWMutex
   - Session storage is thread-safe

6. Configuration Options Pattern
   type ClientOption func(*Client)
   func WithTimeout(timeout time.Duration) ClientOption

================================================================================
END OF CONVERSION PLAN DOCUMENTATION
================================================================================
*/

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// RawMessage 原始消息结构，用于接收微信服务器推送的XML消息
type RawMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MsgID        int64    `xml:"MsgId,omitempty"`
	Event        string   `xml:"Event,omitempty"`
	// 事件相关字段
	EventKey  string `xml:"EventKey,omitempty"`
	Ticket    string `xml:"Ticket,omitempty"`
	Latitude  string `xml:"Latitude,omitempty"`
	Longitude string `xml:"Longitude,omitempty"`
	Precision string `xml:"Precision,omitempty"`
	// 位置信息
	LocationX string `xml:"Location_X,omitempty"`
	LocationY string `xml:"Location_Y,omitempty"`
	Scale     int    `xml:"Scale,omitempty"`
	Label     string `xml:"Label,omitempty"`
	// 媒体信息
	MediaID      string `xml:"MediaId,omitempty"`
	PicURL       string `xml:"PicUrl,omitempty"`
	ThumbMediaID string `xml:"ThumbMediaId,omitempty"`
	// 链接信息
	Title       string `xml:"Title,omitempty"`
	Description string `xml:"Description,omitempty"`
	URL         string `xml:"Url,omitempty"`
	// 语音识别
	Recognition string `xml:"Recognition,omitempty"`
	// 文本消息
	Content string `xml:"Content,omitempty"`
	// 语音消息
	Format string `xml:"Format,omitempty"`
	// 小程序
	PagePath string `xml:"PagePath,omitempty"`
	AppID    string `xml:"AppId,omitempty"`
	ThumbURL string `xml:"ThumbUrl,omitempty"`
	// 群发事件
	Status      string `xml:"Status,omitempty"`
	TotalCount  int    `xml:"TotalCount,omitempty"`
	FilterCount int    `xml:"FilterCount,omitempty"`
	SentCount   int    `xml:"SentCount,omitempty"`
	ErrorCount  int    `xml:"ErrorCount,omitempty"`
	// 模板消息
	TemplateID  string `xml:"TemplateID,omitempty"`
	ClientMsgID string `xml:"ClientMsgId,omitempty"`
}

// Note: MessageType and EventType constants are defined in messages.go and events.go
// This parser uses the constants from those packages for consistency

// ParseError 解析错误类型
type ParseError struct {
	RawData []byte
	Err     error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse message error: %v", e.Err)
}

// ParseMessage 解析微信服务器推送的XML消息
// 这是Parser模块的核心方法，负责将微信服务器的XML消息解析为对应的结构体
//
// 参数:
//   - data: 从微信服务器接收的XML数据
//
// 返回值:
//   - interface{}: 解析后的消息或事件结构体
//   - error: 解析错误
//
// 支持的消息类型:
//   - 文本消息 (TextMessage)
//   - 图片消息 (ImageMessage)
//   - 语音消息 (VoiceMessage)
//   - 视频消息 (VideoMessage/ShortVideoMessage)
//   - 位置消息 (LocationMessage)
//   - 链接消息 (LinkMessage)
//   - 小程序页面消息 (MiniProgramPageMessage)
//
// 支持的事件类型:
//   - 关注/取消关注事件 (SubscribeEvent/UnsubscribeEvent)
//   - 扫码事件 (ScanEvent)
//   - 位置事件 (LocationEvent)
//   - 点击事件 (ClickEvent)
//   - 跳转事件 (ViewEvent)
//   - 群发任务完成事件 (MassSendJobFinishEvent)
//   - 模板消息发送完成事件 (TemplateSendJobFinishEvent)
func ParseMessage(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty message data")
	}

	// 先解析基础消息获取类型
	var raw RawMessage
	if err := xml.Unmarshal(data, &raw); err != nil {
		return nil, &ParseError{RawData: data, Err: err}
	}

	msgType := MessageType(strings.ToLower(raw.MsgType))

	// 处理事件类型
	if msgType == MsgTypeEvent || raw.Event != "" {
		return parseEvent(data, &raw)
	}

	// 处理普通消息
	return parseNormalMessage(data, msgType)
}

// parseEvent 解析事件消息
func parseEvent(data []byte, raw *RawMessage) (interface{}, error) {
	eventType := EventType(raw.Event)

	// 使用已有的raw数据创建事件对象
	switch eventType {
	case EventSubscribe:
		event := &SubscribeEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			EventKey: raw.EventKey,
			Ticket:   raw.Ticket,
		}
		return event, nil

	case EventUnsubscribe:
		event := &UnsubscribeEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
		}
		return event, nil

	case EventScan:
		event := &ScanEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			EventKey: raw.EventKey,
			Ticket:   raw.Ticket,
		}
		return event, nil

	case EventLocation:
		event := &LocationEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			Latitude:  parseFloat64(raw.Latitude),
			Longitude: parseFloat64(raw.Longitude),
			Precision: parseFloat64(raw.Precision),
		}
		return event, nil

	case EventClick:
		event := &ClickEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			EventKey: raw.EventKey,
		}
		return event, nil

	case EventView:
		event := &ViewEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			EventKey: raw.EventKey,
		}
		return event, nil

	case EventMassSendJobFinish:
		event := &MassSendJobFinishEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			Status:      raw.Status,
			TotalCount:  raw.TotalCount,
			FilterCount: raw.FilterCount,
			SentCount:   raw.SentCount,
			ErrorCount:  raw.ErrorCount,
		}
		return event, nil

	case EventTemplateSendJobFinish:
		event := &TemplateSendJobFinishEvent{
			BaseEvent: BaseEvent{
				BaseMessage: BaseMessage{
					ToUserName:   raw.ToUserName,
					FromUserName: raw.FromUserName,
					CreateTime:   raw.CreateTime,
					MsgType:      "event",
				},
				Event: raw.Event,
			},
			MsgID:  raw.MsgID,
			Status: raw.Status,
		}
		return event, nil

	default:
		// 未知事件类型，返回基础事件结构
		event := &BaseEvent{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "event",
			},
			Event: raw.Event,
		}
		return event, nil
	}
}

// parseNormalMessage 解析普通消息
func parseNormalMessage(data []byte, msgType MessageType) (interface{}, error) {
	// 重新解析原始数据以获取完整信息
	var raw RawMessage
	if err := xml.Unmarshal(data, &raw); err != nil {
		return nil, &ParseError{RawData: data, Err: err}
	}

	switch msgType {
	case MsgTypeText:
		msg := &TextMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "text",
				MsgID:        raw.MsgID,
			},
			Content: raw.Content,
		}
		return msg, nil

	case MsgTypeImage:
		msg := &ImageMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "image",
				MsgID:        raw.MsgID,
			},
			PicURL:  raw.PicURL,
			MediaID: raw.MediaID,
		}
		return msg, nil

	case MsgTypeVoice:
		msg := &VoiceMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "voice",
				MsgID:        raw.MsgID,
			},
			MediaID:     raw.MediaID,
			Format:      raw.Format,
			Recognition: raw.Recognition,
		}
		return msg, nil

	case MsgTypeVideo:
		msg := &VideoMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "video",
				MsgID:        raw.MsgID,
			},
			MediaID:      raw.MediaID,
			ThumbMediaID: raw.ThumbMediaID,
		}
		return msg, nil

	case MsgTypeShortVideo:
		msg := &ShortVideoMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "shortvideo",
				MsgID:        raw.MsgID,
			},
			MediaID:      raw.MediaID,
			ThumbMediaID: raw.ThumbMediaID,
		}
		return msg, nil

	case MsgTypeLocation:
		msg := &LocationMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "location",
				MsgID:        raw.MsgID,
			},
			LocationX: parseFloat64(raw.LocationX),
			LocationY: parseFloat64(raw.LocationY),
			Scale:     raw.Scale,
			Label:     raw.Label,
		}
		return msg, nil

	case MsgTypeLink:
		msg := &LinkMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "link",
				MsgID:        raw.MsgID,
			},
			Title:       raw.Title,
			Description: raw.Description,
			URL:         raw.URL,
		}
		return msg, nil

	case MsgTypeMiniProgramPage:
		msg := &MiniProgramPageMessage{
			BaseMessage: BaseMessage{
				ToUserName:   raw.ToUserName,
				FromUserName: raw.FromUserName,
				CreateTime:   raw.CreateTime,
				MsgType:      "miniprogrampage",
				MsgID:        raw.MsgID,
			},
			AppID:        raw.AppID,
			Title:        raw.Title,
			PagePath:     raw.PagePath,
			ThumbURL:     raw.ThumbURL,
			ThumbMediaID: raw.ThumbMediaID,
		}
		return msg, nil

	default:
		// 未知消息类型，返回基础消息结构
		msg := &BaseMessage{
			ToUserName:   raw.ToUserName,
			FromUserName: raw.FromUserName,
			CreateTime:   raw.CreateTime,
			MsgType:      string(msgType),
			MsgID:        raw.MsgID,
		}
		return msg, nil
	}
}

// 辅助函数：解析浮点数
func parseFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

// MessageParser 消息解析器接口
// 可以用于实现不同的解析策略或模拟测试
type MessageParser interface {
	Parse(data []byte) (interface{}, error)
}

// DefaultParser 默认消息解析器
type DefaultParser struct{}

// Parse 实现MessageParser接口
func (p *DefaultParser) Parse(data []byte) (interface{}, error) {
	return ParseMessage(data)
}

// NewDefaultParser 创建默认解析器实例
func NewDefaultParser() MessageParser {
	return &DefaultParser{}
}
