package wechatgo

import (
	"encoding/xml"
	"fmt"
	"time"
)

// Reply 回复接口
type Reply interface {
	Render() ([]byte, error)
}

// BaseReply 基础回复
type BaseReply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

// NewBaseReply 创建基础回复
func NewBaseReply(toUser, fromUser, msgType string) BaseReply {
	return BaseReply{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      msgType,
	}
}

// TextReply 文本回复
type TextReply struct {
	BaseReply
	Content string `xml:"Content"`
}

// NewTextReply 创建文本回复
func NewTextReply(toUser, fromUser, content string) *TextReply {
	return &TextReply{
		BaseReply: NewBaseReply(toUser, fromUser, "text"),
		Content:   content,
	}
}

// Render 渲染为 XML
func (r *TextReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// ImageReply 图片回复
type ImageReply struct {
	BaseReply
	Image struct {
		MediaID string `xml:"MediaId"`
	} `xml:"Image"`
}

// NewImageReply 创建图片回复
func NewImageReply(toUser, fromUser, mediaID string) *ImageReply {
	reply := &ImageReply{
		BaseReply: NewBaseReply(toUser, fromUser, "image"),
	}
	reply.Image.MediaID = mediaID
	return reply
}

// Render 渲染为 XML
func (r *ImageReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// VoiceReply 语音回复
type VoiceReply struct {
	BaseReply
	Voice struct {
		MediaID string `xml:"MediaId"`
	} `xml:"Voice"`
}

// NewVoiceReply 创建语音回复
func NewVoiceReply(toUser, fromUser, mediaID string) *VoiceReply {
	reply := &VoiceReply{
		BaseReply: NewBaseReply(toUser, fromUser, "voice"),
	}
	reply.Voice.MediaID = mediaID
	return reply
}

// Render 渲染为 XML
func (r *VoiceReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// VideoReply 视频回复
type VideoReply struct {
	BaseReply
	Video struct {
		MediaID     string `xml:"MediaId"`
		Title       string `xml:"Title,omitempty"`
		Description string `xml:"Description,omitempty"`
	} `xml:"Video"`
}

// NewVideoReply 创建视频回复
func NewVideoReply(toUser, fromUser, mediaID, title, description string) *VideoReply {
	reply := &VideoReply{
		BaseReply: NewBaseReply(toUser, fromUser, "video"),
	}
	reply.Video.MediaID = mediaID
	reply.Video.Title = title
	reply.Video.Description = description
	return reply
}

// Render 渲染为 XML
func (r *VideoReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// MusicReply 音乐回复
type MusicReply struct {
	BaseReply
	Music struct {
		Title        string `xml:"Title,omitempty"`
		Description  string `xml:"Description,omitempty"`
		MusicURL     string `xml:"MusicUrl,omitempty"`
		HQMusicURL   string `xml:"HQMusicUrl,omitempty"`
		ThumbMediaID string `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// NewMusicReply 创建音乐回复
func NewMusicReply(toUser, fromUser, thumbMediaID string) *MusicReply {
	reply := &MusicReply{
		BaseReply: NewBaseReply(toUser, fromUser, "music"),
	}
	reply.Music.ThumbMediaID = thumbMediaID
	return reply
}

// Render 渲染为 XML
func (r *MusicReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// Article 图文消息文章
type Article struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicURL      string `xml:"PicUrl"`
	URL         string `xml:"Url"`
}

// NewsReply 图文消息回复
type NewsReply struct {
	BaseReply
	ArticleCount int       `xml:"ArticleCount"`
	Articles     []Article `xml:"Articles>item"`
}

// NewNewsReply 创建图文消息回复
func NewNewsReply(toUser, fromUser string, articles []Article) *NewsReply {
	return &NewsReply{
		BaseReply:    NewBaseReply(toUser, fromUser, "news"),
		ArticleCount: len(articles),
		Articles:     articles,
	}
}

// Render 渲染为 XML
func (r *NewsReply) Render() ([]byte, error) {
	return xml.Marshal(r)
}

// CreateReply 根据消息创建回复
func CreateReply(msg Message, replyType, content string) (Reply, error) {
	toUser := msg.GetFromUserName()
	fromUser := msg.GetToUserName()

	switch replyType {
	case "text":
		return NewTextReply(toUser, fromUser, content), nil
	default:
		return nil, fmt.Errorf("unsupported reply type: %s", replyType)
	}
}
