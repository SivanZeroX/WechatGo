package pay

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/wechatpy/wechatgo/pay/api"
)

// Client 微信支付客户端
type Client struct {
	AppID      string `json:"appid"`     // 公众号APPID
	APIKey     string `json:"api_key"`   // 商户API密钥
	MchID      string `json:"mch_id"`    // 商户号
	CertPath   string `json:"cert_path"` // 商户证书路径
	KeyPath    string `json:"key_path"`  // 商户私钥路径
	httpClient api.HTTPClient

	// API 模块
	Order       *api.OrderAPI       `json:"-"` // 订单接口
	Refund      *api.RefundAPI      `json:"-"` // 退款接口
	JsAPI       *api.JsAPIAPI       `json:"-"` // JSAPI接口
	MicroPay    *api.MicroPayAPI    `json:"-"` // 刷卡支付接口
	Tools       *api.ToolsAPI       `json:"-"` // 工具类接口
	RedPack     *api.RedPackAPI     `json:"-"` // 红包接口
	Transfer    *api.TransferAPI    `json:"-"` // 企业付款接口
	Coupon      *api.CouponAPI      `json:"-"` // 代金券接口
	ProfitShare *api.ProfitShareAPI `json:"-"` // 分账接口
}

// NewClient 创建微信支付客户端
func NewClient(appID, apiKey, mchID, certPath, keyPath string, httpClient api.HTTPClient) *Client {
	c := &Client{
		AppID:      appID,
		APIKey:     apiKey,
		MchID:      mchID,
		CertPath:   certPath,
		KeyPath:    keyPath,
		httpClient: httpClient,
	}

	// 初始化 API 模块
	c.Order = api.NewOrderAPI(c)
	c.Refund = api.NewRefundAPI(c)
	c.JsAPI = api.NewJsAPIAPI(c)
	c.MicroPay = api.NewMicroPayAPI(c)
	c.Tools = api.NewToolsAPI(c)
	c.RedPack = api.NewRedPackAPI(c)
	c.Transfer = api.NewTransferAPI(c)
	c.Coupon = api.NewCouponAPI(c)
	c.ProfitShare = api.NewProfitShareAPI(c)

	return c
}

// APIBaseURL 微信支付API基础URL
const APIBaseURL = "https://api.mch.weixin.qq.com/"

// GetAppID 返回AppID
func (c *Client) GetAppID() string {
	return c.AppID
}

// GetMchID 返回商户号
func (c *Client) GetMchID() string {
	return c.MchID
}

// GetAPIKey 返回API密钥
func (c *Client) GetAPIKey() string {
	return c.APIKey
}

// GetHTTPClient 返回HTTP客户端
func (c *Client) GetHTTPClient() api.HTTPClient {
	return c.httpClient
}

// Get implements api.HTTPClient
func (c *Client) Get(url string) (*http.Response, error) {
	return c.httpClient.Get(url)
}

// Post implements api.HTTPClient
func (c *Client) Post(url string, data []byte, headers map[string]string) (*http.Response, error) {
	return c.httpClient.Post(url, data, headers)
}

// GetPrepayID 获取预支付ID
func (c *Client) GetPrepayID(req *api.PrepayRequest) (string, error) {
	result, err := c.Order.GetPrepayID(req)
	if err != nil {
		return "", err
	}
	return result.PrepayID, nil
}

// GenerateJSAPIPayParams 生成JSAPI支付参数
func (c *Client) GenerateJSAPIPayParams(prepayID string) (map[string]string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonceStr := api.RandomString(32)
	pkg := fmt.Sprintf("prepay_id=%s", prepayID)

	// 生成签名
	signStr := fmt.Sprintf("appId=%s&nonceStr=%s&package=%s&signType=MD5&timeStamp=%s&key=%s",
		c.AppID, nonceStr, pkg, timestamp, c.APIKey)
	md5Hash := api.MD5(signStr)
	sign := strings.ToUpper(md5Hash)

	return map[string]string{
		"appId":     c.AppID,
		"timeStamp": timestamp,
		"nonceStr":  nonceStr,
		"package":   pkg,
		"signType":  "MD5",
		"paySign":   sign,
	}, nil
}

// VerifySignature 验证签名
func (c *Client) VerifySignature(params map[string]string, sign string) bool {
	// 生成签名
	delete(params, "sign")
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	signStr := ""
	for _, k := range keys {
		signStr += fmt.Sprintf("%s=%s&", k, params[k])
	}
	signStr += fmt.Sprintf("key=%s", c.APIKey)

	// 计算签名
	md5Hash := api.MD5(signStr)
	calcSign := strings.ToUpper(md5Hash)

	return calcSign == sign
}
