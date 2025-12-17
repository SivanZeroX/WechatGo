package wechatgo

// EventType 事件类型
type EventType string

const (
	EventSubscribe             EventType = "subscribe"
	EventUnsubscribe           EventType = "unsubscribe"
	EventScan                  EventType = "SCAN"
	EventLocation              EventType = "LOCATION"
	EventClick                 EventType = "CLICK"
	EventView                  EventType = "VIEW"
	EventMassSendJobFinish     EventType = "MASSSENDJOBFINISH"
	EventTemplateSendJobFinish EventType = "TEMPLATESENDJOBFINISH"
)

// BaseEvent 基础事件
type BaseEvent struct {
	BaseMessage
	Event string `xml:"Event"`
}

// SubscribeEvent 关注事件
type SubscribeEvent struct {
	BaseEvent
	EventKey string `xml:"EventKey,omitempty"`
	Ticket   string `xml:"Ticket,omitempty"`
}

// UnsubscribeEvent 取消关注事件
type UnsubscribeEvent struct {
	BaseEvent
}

// ScanEvent 扫描二维码事件
type ScanEvent struct {
	BaseEvent
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

// LocationEvent 上报地理位置事件
type LocationEvent struct {
	BaseEvent
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision float64 `xml:"Precision"`
}

// ClickEvent 点击菜单拉取消息事件
type ClickEvent struct {
	BaseEvent
	EventKey string `xml:"EventKey"`
}

// ViewEvent 点击菜单跳转链接事件
type ViewEvent struct {
	BaseEvent
	EventKey string `xml:"EventKey"`
}

// MassSendJobFinishEvent 群发消息任务完成事件
type MassSendJobFinishEvent struct {
	BaseEvent
	MsgID       int64  `xml:"MsgID"`
	Status      string `xml:"Status"`
	TotalCount  int    `xml:"TotalCount"`
	FilterCount int    `xml:"FilterCount"`
	SentCount   int    `xml:"SentCount"`
	ErrorCount  int    `xml:"ErrorCount"`
}

// TemplateSendJobFinishEvent 模板消息任务完成事件
type TemplateSendJobFinishEvent struct {
	BaseEvent
	MsgID  int64  `xml:"MsgID"`
	Status string `xml:"Status"`
}
