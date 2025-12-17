package api

import "fmt"

// MessageAPI 消息发送 API
type MessageAPI struct {
	*BaseAPI
}

// NewMessageAPI 创建消息发送 API
func NewMessageAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MessageAPI {
	return &MessageAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// SendText 发送文本消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
func (api *MessageAPI) SendText(openID, content string, kfAccount string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"touser":  openID,
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
	}

	if kfAccount != "" {
		data["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}

	return api.Post("/message/custom/send", data)
}

// SendImage 发送图片消息
func (api *MessageAPI) SendImage(openID, mediaID string, kfAccount string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"touser":  openID,
		"msgtype": "image",
		"image": map[string]string{
			"media_id": mediaID,
		},
	}

	if kfAccount != "" {
		data["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}

	return api.Post("/message/custom/send", data)
}

// SendVoice 发送语音消息
func (api *MessageAPI) SendVoice(openID, mediaID string, kfAccount string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"touser":  openID,
		"msgtype": "voice",
		"voice": map[string]string{
			"media_id": mediaID,
		},
	}

	if kfAccount != "" {
		data["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}

	return api.Post("/message/custom/send", data)
}

// SendVideo 发送视频消息
func (api *MessageAPI) SendVideo(openID, mediaID, thumbMediaID, title, description string, kfAccount string) (map[string]interface{}, error) {
	video := map[string]string{
		"media_id": mediaID,
	}
	if thumbMediaID != "" {
		video["thumb_media_id"] = thumbMediaID
	}
	if title != "" {
		video["title"] = title
	}
	if description != "" {
		video["description"] = description
	}

	data := map[string]interface{}{
		"touser":  openID,
		"msgtype": "video",
		"video":   video,
	}

	if kfAccount != "" {
		data["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}

	return api.Post("/message/custom/send", data)
}

// NewsArticle 图文消息文章
type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// SendNews 发送图文消息
func (api *MessageAPI) SendNews(openID string, articles []NewsArticle, kfAccount string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"touser":  openID,
		"msgtype": "news",
		"news": map[string]interface{}{
			"articles": articles,
		},
	}

	if kfAccount != "" {
		data["customservice"] = map[string]string{
			"kf_account": kfAccount,
		}
	}

	return api.Post("/message/custom/send", data)
}

// DeleteMass 删除群发消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) DeleteMass(msgID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"msg_id": msgID,
	}
	return api.Post("/message/mass/delete", data)
}

// sendMassMessage 发送群发消息的内部方法
func (api *MessageAPI) sendMassMessage(tagOrUsers interface{}, msgType string, msg map[string]interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"msgtype":             msgType,
		"send_ignore_reprint": sendIgnoreReprint,
	}
	if clientMsgID != nil {
		data["clientmsgid"] = *clientMsgID
	}

	var endpoint string
	if !preview {
		switch v := tagOrUsers.(type) {
		case []string:
			// 按 OpenID 列表群发
			data["touser"] = v
			endpoint = "/message/mass/send"
		case int:
			// 按标签群发
			data["filter"] = map[string]interface{}{
				"tag_id":    v,
				"is_to_all": isToAll,
			}
			endpoint = "/message/mass/sendall"
		default:
			if isToAll {
				// 发送给全部用户
				data["filter"] = map[string]interface{}{
					"is_to_all": true,
				}
				endpoint = "/message/mass/sendall"
			} else {
				return nil, fmt.Errorf("invalid tag_or_users type")
			}
		}
	} else {
		// 预览接口
		if openID, ok := tagOrUsers.(string); ok {
			data["touser"] = openID
			endpoint = "/message/mass/preview"
		} else {
			return nil, fmt.Errorf("preview mode requires string openid")
		}
	}

	data = mergeMaps(data, msg)
	return api.Post(endpoint, data)
}

// SendMassText 群发文本消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) SendMassText(content string, tagOrUsers interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	msg := map[string]interface{}{
		"text": map[string]string{
			"content": content,
		},
	}
	return api.sendMassMessage(tagOrUsers, "text", msg, isToAll, preview, sendIgnoreReprint, clientMsgID)
}

// SendMassImage 群发图片消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) SendMassImage(mediaID string, tagOrUsers interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	msg := map[string]interface{}{
		"image": map[string]string{
			"media_id": mediaID,
		},
	}
	return api.sendMassMessage(tagOrUsers, "image", msg, isToAll, preview, sendIgnoreReprint, clientMsgID)
}

// SendMassVoice 群发语音消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) SendMassVoice(mediaID string, tagOrUsers interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	msg := map[string]interface{}{
		"voice": map[string]string{
			"media_id": mediaID,
		},
	}
	return api.sendMassMessage(tagOrUsers, "voice", msg, isToAll, preview, sendIgnoreReprint, clientMsgID)
}

// SendMassVideo 群发视频消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) SendMassVideo(mediaID string, tagOrUsers interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	msg := map[string]interface{}{
		"mpvideo": map[string]string{
			"media_id": mediaID,
		},
	}
	return api.sendMassMessage(tagOrUsers, "mpvideo", msg, isToAll, preview, sendIgnoreReprint, clientMsgID)
}

// SendMassNews 群发图文消息
// https://mp.weixin.qq.com/wiki?id=mp1481187827_i0l21
func (api *MessageAPI) SendMassNews(mediaID string, tagOrUsers interface{}, isToAll, preview bool, sendIgnoreReprint int, clientMsgID *string) (map[string]interface{}, error) {
	msg := map[string]interface{}{
		"mpnews": map[string]string{
			"media_id": mediaID,
		},
	}
	return api.sendMassMessage(tagOrUsers, "mpnews", msg, isToAll, preview, sendIgnoreReprint, clientMsgID)
}

// mergeMaps 合并两个 map
func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}
