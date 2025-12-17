package api

import (
	"fmt"
)

// MediaAPI 媒体相关API
type MediaAPI struct {
	*BaseAPI
}

// NewMediaAPI 创建媒体API
func NewMediaAPI(client *Client) *MediaAPI {
	return &MediaAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// UploadImage 上传图片
func (api *MediaAPI) UploadImage(imageData []byte) (*UploadImageResponse, error) {
	// TODO: 实现上传图片逻辑
	return nil, fmt.Errorf("not implemented")
}

// UploadImageResponse 上传图片响应
type UploadImageResponse struct {
	ImageURL string `json:"image_url"` // 图片URL
	MediaID  string `json:"media_id"`  // 媒体ID
}
