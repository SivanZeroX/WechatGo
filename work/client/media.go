package client

import (
	"fmt"
	"net/http"
)

// MediaAPI 媒体管理API
type MediaAPI struct {
	BaseAPI interface {
		Get(url string, params map[string]string) (map[string]interface{}, error)
		Post(url string, data interface{}) (map[string]interface{}, error)
		GetAccessToken() (string, error)
	}
}

// NewMediaAPI 创建媒体管理API
func NewMediaAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MediaAPI {
	return &MediaAPI{
		BaseAPI: client,
	}
}

// Upload 上传临时素材
// https://developer.work.weixin.qq.com/document/path/90253
func (api *MediaAPI) Upload(mediaType, filePath string) (*UploadMediaResponse, error) {
	// TODO: 实现文件上传逻辑
	return nil, fmt.Errorf("not implemented")
}

// Get 获取临时素材
// https://developer.work.weixin.qq.com/document/path/90253
func (api *MediaAPI) Get(mediaID string) (*http.Response, error) {
	// TODO: 实现获取临时素材逻辑
	// 需要使用原始HTTP请求而不是JSON API
	return nil, fmt.Errorf("not implemented")
}

// UploadImg 上传图片
// https://developer.work.weixin.qq.com/document/path/90253
func (api *MediaAPI) UploadImg(filePath string) (*UploadImgResponse, error) {
	// TODO: 实现图片上传逻辑
	return nil, fmt.Errorf("not implemented")
}

// ==================== 响应结构体 ====================

// UploadMediaResponse 上传临时素材响应
type UploadMediaResponse struct {
	MediaID   string `json:"media_id"`   // 媒体ID
	CreatedAt int64  `json:"created_at"` // 创建时间
}

// UploadImgResponse 上传图片响应
type UploadImgResponse struct {
	ImageURL string `json:"image_url"` // 图片URL
}
