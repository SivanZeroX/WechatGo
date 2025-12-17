package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/wechatpy/wechatgo"
)

// CustomServiceAPI 客服消息管理 API
type CustomServiceAPI struct {
	*BaseAPI
}

// NewCustomServiceAPI 创建客服消息 API
func NewCustomServiceAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	Upload(url string, fileName string, file io.Reader) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *CustomServiceAPI {
	return &CustomServiceAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// AddAccount 添加客服账号
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#添加客服账号
func (api *CustomServiceAPI) AddAccount(account, nickname, password string) (map[string]interface{}, error) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	data := map[string]interface{}{
		"kf_account": account,
		"nickname":   nickname,
		"password":   password,
	}
	return api.Post("/customservice/kfaccount/add", data)
}

// UpdateAccount 修改客服账号
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#修改客服账号
func (api *CustomServiceAPI) UpdateAccount(account, nickname, password string) (map[string]interface{}, error) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	data := map[string]interface{}{
		"kf_account": account,
		"nickname":   nickname,
		"password":   password,
	}
	return api.Post("/customservice/kfaccount/update", data)
}

// DeleteAccount 删除客服账号
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#删除客服账号
func (api *CustomServiceAPI) DeleteAccount(account string) (map[string]interface{}, error) {
	return api.Get("/customservice/kfaccount/del", map[string]string{"kf_account": account})
}

// Account 客服账号信息
type Account struct {
	Account      string `json:"kf_account"`
	Nickname     string `json:"kf_nick"`
	AccountID    int    `json:"kf_id"`
	HeadImgURL   string `json:"head_img_url"`
	InviteWX     string `json:"invite_wx"`
	InviteExpire int    `json:"invite_expire_time"`
	InviteStatus int    `json:"invite_status"`
}

// GetAccounts 获取所有客服账号
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#获取所有客服账号
func (api *CustomServiceAPI) GetAccounts() ([]Account, error) {
	result, err := api.Get("/customservice/getkflist", nil)
	if err != nil {
		return nil, err
	}

	if kfList, ok := result["kf_list"].([]interface{}); ok {
		accounts := make([]Account, len(kfList))
		for i, item := range kfList {
			if accountMap, ok := item.(map[string]interface{}); ok {
				accounts[i] = Account{
					Account:      accountMap["kf_account"].(string),
					Nickname:     accountMap["kf_nick"].(string),
					AccountID:    int(accountMap["kf_id"].(float64)),
					HeadImgURL:   accountMap["head_img_url"].(string),
					InviteWX:     accountMap["invite_wx"].(string),
					InviteExpire: int(accountMap["invite_expire_time"].(float64)),
					InviteStatus: int(accountMap["invite_status"].(float64)),
				}
			}
		}
		return accounts, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// UploadHeadImg 设置客服账号的头像
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html#设置客服账号的头像
func (api *CustomServiceAPI) UploadHeadImg(account, filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return api.uploadHeadImg(account, filePath, file)
}

// UploadHeadImgReader 设置客服账号的头像（从Reader）
func (api *CustomServiceAPI) UploadHeadImgReader(account, fileName string, reader io.Reader) (map[string]interface{}, error) {
	return api.uploadHeadImg(account, fileName, reader)
}

// uploadHeadImg 内部上传头像方法
func (api *CustomServiceAPI) uploadHeadImg(account, fileName string, reader io.Reader) (map[string]interface{}, error) {
	token, err := api.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s", token, account)

	resp, err := api.httpClient.Post(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

// OnlineAccount 在线客服信息
type OnlineAccount struct {
	Account       string `json:"kf_account"`
	Nickname      string `json:"kf_nick"`
	Status        int    `json:"status"`
	KFID          int    `json:"kf_id"`
	WorkTime      string `json:"work_time"`
	CustomerCount int    `json:"customer_count"`
}

// GetOnlineAccounts 获取在线客服接待信息
// http://mp.weixin.qq.com/wiki/9/6fff6f191ef92c126b043ada035cc935.html
func (api *CustomServiceAPI) GetOnlineAccounts() ([]OnlineAccount, error) {
	result, err := api.Get("/customservice/getonlinekflist", nil)
	if err != nil {
		return nil, err
	}

	if kfList, ok := result["kf_online_list"].([]interface{}); ok {
		accounts := make([]OnlineAccount, len(kfList))
		for i, item := range kfList {
			if accountMap, ok := item.(map[string]interface{}); ok {
				accounts[i] = OnlineAccount{
					Account:       accountMap["kf_account"].(string),
					Nickname:      accountMap["kf_nick"].(string),
					Status:        int(accountMap["status"].(float64)),
					KFID:          int(accountMap["kf_id"].(float64)),
					WorkTime:      accountMap["work_time"].(string),
					CustomerCount: int(accountMap["customer_count"].(float64)),
				}
			}
		}
		return accounts, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// CreateSession 多客服创建会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (api *CustomServiceAPI) CreateSession(openid, account, text string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"openid":     openid,
		"kf_account": account,
	}
	if text != "" {
		data["text"] = text
	}
	return api.Post("/customservice/kfsession/create", data)
}

// CloseSession 多客服关闭会话
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (api *CustomServiceAPI) CloseSession(openid, account, text string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"openid":     openid,
		"kf_account": account,
	}
	if text != "" {
		data["text"] = text
	}
	return api.Post("/customservice/kfsession/close", data)
}

// Session 客户会话状态
type Session struct {
	SessionCreateTime int    `json:"createtime"`
	KFAccount         string `json:"kf_account"`
	KFID              int    `json:"kf_id"`
	KFNickname        string `json:"kf_nick"`
	State             int    `json:"state"`
	WaitTime          int    `json:"waitcaselist"`
}

// GetSession 获取客户的会话状态
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (api *CustomServiceAPI) GetSession(openid string) (map[string]interface{}, error) {
	return api.Get("/customservice/kfsession/getsession", map[string]string{"openid": openid})
}

// GetSessionList 获取客服的会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (api *CustomServiceAPI) GetSessionList(account string) (map[string]interface{}, error) {
	return api.Get("/customservice/kfsession/getsessionlist", map[string]string{"kf_account": account})
}

// GetWaitCase 获取未接入会话列表
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html
func (api *CustomServiceAPI) GetWaitCase() (map[string]interface{}, error) {
	return api.Get("/customservice/kfsession/getwaitcase", nil)
}

// GetRecords 获取客服聊天记录
// https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html
func (api *CustomServiceAPI) GetRecords(startTime, endTime int64, msgID, number int) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"starttime": startTime,
		"endtime":   endTime,
		"msgid":     msgID,
		"number":    number,
	}
	return api.Post("/customservice/msgrecord/getmsglist", data)
}
