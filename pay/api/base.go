package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// BaseAPI 基础API
type BaseAPI struct {
	client Client
}

// NewBaseAPI 创建基础API
func NewBaseAPI(client Client) *BaseAPI {
	return &BaseAPI{
		client: client,
	}
}

// Client HTTP客户端接口
type Client interface {
	HTTPClient
	GetAppID() string
	GetMchID() string
	GetAPIKey() string
	GetHTTPClient() HTTPClient
	GenerateJSAPIPayParams(prepayID string) (map[string]string, error)
}

// HTTPClient HTTP客户端接口
type HTTPClient interface {
	Post(url string, data []byte, headers map[string]string) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

// BaseResponse 基础响应
type BaseResponse struct {
	ReturnCode string `json:"return_code"`  // 返回状态码
	ReturnMsg  string `json:"return_msg"`   // 返回信息
	AppID      string `json:"appid"`        // 公众账号ID
	MchID      string `json:"mch_id"`       // 商户号
	DeviceInfo string `json:"device_info"`  // 设备号
	NonceStr   string `json:"nonce_str"`    // 随机字符串
	Sign       string `json:"sign"`         // 签名
	ResultCode string `json:"result_code"`  // 业务结果码
	ErrCode    string `json:"err_code"`     // 错误代码
	ErrCodeDes string `json:"err_code_des"` // 错误代码描述
}

// PrepayRequest 预支付请求
type PrepayRequest struct {
	AppID          string `json:"appid"`            // 公众号ID
	MchID          string `json:"mch_id"`           // 商户号
	Body           string `json:"body"`             // 商品描述
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TotalFee       int    `json:"total_fee"`        // 总金额（分）
	SpbillCreateIP string `json:"spbill_create_ip"` // 终端IP
	NotifyURL      string `json:"notify_url"`       // 通知地址
	TradeType      string `json:"trade_type"`       // 交易类型
	OpenID         string `json:"openid,omitempty"` // 用户标识（JSAPI支付必填）
}

// PrepayResponse 预支付响应
type PrepayResponse struct {
	BaseResponse
	AppID      string `json:"appid"`       // 公众账号ID
	MchID      string `json:"mch_id"`      // 商户号
	DeviceInfo string `json:"device_info"` // 设备号
	NonceStr   string `json:"nonce_str"`   // 随机字符串
	Sign       string `json:"sign"`        // 签名
	TradeType  string `json:"trade_type"`  // 交易类型
	PrepayID   string `json:"prepay_id"`   // 预支付交易会话标识
	CodeURL    string `json:"code_url"`    // 二维码链接（NATIVE支付）
}

// APIBaseURL 微信支付API基础URL
const APIBaseURL = "https://api.mch.weixin.qq.com/"

// RandomString 生成随机字符串
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[i%len(charset)]
	}
	return string(result)
}

// MD5 计算MD5
func MD5(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateSignature 生成签名
func GenerateSignature(params map[string]string, key string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signStr := ""
	for _, k := range keys {
		if params[k] != "" {
			signStr += fmt.Sprintf("%s=%s&", k, params[k])
		}
	}
	signStr += fmt.Sprintf("key=%s", key)

	// 计算签名
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(signStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
