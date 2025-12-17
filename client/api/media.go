package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/wechatpy/wechatgo"
)

// MediaAPI 素材管理 API
type MediaAPI struct {
	*BaseAPI
}

// NewMediaAPI 创建素材管理 API
func NewMediaAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	Upload(url string, fileName string, file io.Reader) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *MediaAPI {
	return &MediaAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// Upload 上传临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func (api *MediaAPI) Upload(mediaType, filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return api.uploadMedia(mediaType, filepath.Base(filePath), file)
}

// UploadReader 上传临时素材（从Reader）
func (api *MediaAPI) UploadReader(mediaType, fileName string, reader io.Reader) (map[string]interface{}, error) {
	return api.uploadMedia(mediaType, fileName, reader)
}

// uploadMedia 内部上传方法
func (api *MediaAPI) uploadMedia(mediaType, fileName string, reader io.Reader) (map[string]interface{}, error) {
	// 构建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加type字段
	writer.WriteField("type", mediaType)

	// 添加文件
	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}
	writer.Close()

	// 构建URL
	token, err := api.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", token, mediaType)

	// 发送请求
	resp, err := api.httpClient.Post(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// 检查错误
	if errcode, ok := result["errcode"]; ok {
		errcodeInt := int(errcode.(float64))
		if errcodeInt != 0 {
			errmsg := ""
			if msg, ok := result["errmsg"]; ok {
				errmsg = msg.(string)
			}
			return nil, wechatgo.NewError(errcodeInt, errmsg)
		}
	}

	return result, nil
}

// Download 获取临时素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html
func (api *MediaAPI) Download(mediaID string) ([]byte, error) {
	token, err := api.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", token, mediaID)

	resp, err := api.httpClient.GetRaw(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// GetURL 获取临时素材下载地址
func (api *MediaAPI) GetURL(mediaID string) string {
	token, _ := api.GetAccessToken()
	return fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s", token, mediaID)
}

// UploadVideo 上传视频（用于群发视频消息）
func (api *MediaAPI) UploadVideo(mediaID, title, description string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"media_id":    mediaID,
		"title":       title,
		"description": description,
	}
	return api.Post("/media/uploadvideo", data)
}

// Article 图文消息文章
type Article struct {
	ThumbMediaID     string `json:"thumb_media_id"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Author           string `json:"author,omitempty"`
	ContentSourceURL string `json:"content_source_url,omitempty"`
	Digest           string `json:"digest,omitempty"`
	ShowCoverPic     int    `json:"show_cover_pic,omitempty"`
}

// UploadArticles 上传图文消息素材
func (api *MediaAPI) UploadArticles(articles []Article) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"articles": articles,
	}
	return api.Post("/media/uploadnews", data)
}

// UploadImage 上传群发消息内的图片
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (api *MediaAPI) UploadImage(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := api.uploadImage(filepath.Base(filePath), file)
	if err != nil {
		return "", err
	}

	// 返回图片URL
	if url, ok := result["url"].(string); ok {
		return url, nil
	}
	return "", fmt.Errorf("unexpected response format")
}

// UploadImageReader 上传群发消息内的图片（从Reader）
func (api *MediaAPI) UploadImageReader(fileName string, reader io.Reader) (string, error) {
	result, err := api.uploadImage(fileName, reader)
	if err != nil {
		return "", err
	}

	// 返回图片URL
	if url, ok := result["url"].(string); ok {
		return url, nil
	}
	return "", fmt.Errorf("unexpected response format")
}

// uploadImage 内部上传图片方法
func (api *MediaAPI) uploadImage(fileName string, reader io.Reader) (map[string]interface{}, error) {
	// 构建multipart表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件
	part, err := writer.CreateFormFile("media", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}
	writer.Close()

	// 构建URL
	token, err := api.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s", token)

	// 发送请求
	resp, err := api.httpClient.Post(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// 检查错误
	if errcode, ok := result["errcode"]; ok {
		errcodeInt := int(errcode.(float64))
		if errcodeInt != 0 {
			errmsg := ""
			if msg, ok := result["errmsg"]; ok {
				errmsg = msg.(string)
			}
			return nil, wechatgo.NewError(errcodeInt, errmsg)
		}
	}

	return result, nil
}
