package api

import (
	"encoding/xml"
	"fmt"
)

// CouponAPI 代金券接口
type CouponAPI struct {
	*BaseAPI
}

// NewCouponAPI 创建代金券接口
func NewCouponAPI(client Client) *CouponAPI {
	return &CouponAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// QueryCoupons 查询代金券
func (api *CouponAPI) QueryCoupons(req *QueryCouponsRequest) (*QueryCouponsResponse, error) {
	// 构建请求参数
	params := map[string]string{
		"appid":     api.client.GetAppID(),
		"mch_id":    api.client.GetMchID(),
		"coupon_id": req.CouponID,
		"openid":    req.OpenID,
		"stock_id":  req.StockID,
		"nonce_str": RandomString(32),
	}

	// 生成签名
	sign := GenerateSignature(params, api.client.GetAPIKey())
	params["sign"] = sign

	// 转换为XML
	xmlData, err := xml.Marshal(map[string]string(params))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal xml: %w", err)
	}

	// 发送请求
	resp, err := api.client.GetHTTPClient().Post(
		APIBaseURL+"marketing/coupon/query",
		xmlData,
		map[string]string{
			"Content-Type": "application/xml",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query coupons: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var result QueryCouponsResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// QueryCouponsRequest 查询代金券请求
type QueryCouponsRequest struct {
	CouponID string `json:"coupon_id"` // 代金券ID
	OpenID   string `json:"openid"`    // 用户的OpenID
	StockID  string `json:"stock_id"`  // 批次号
}

// QueryCouponsResponse 查询代金券响应
type QueryCouponsResponse struct {
	BaseResponse
	TotalNum   int          `json:"total_num"`   // 代金券总数
	CouponList []CouponInfo `json:"coupon_list"` // 代金券列表
}

// CouponInfo 代金券信息
type CouponInfo struct {
	CouponID           string `json:"coupon_id"`            // 代金券ID
	CouponType         string `json:"coupon_type"`          // 代金券类型
	CouponState        string `json:"coupon_state"`         // 代金券状态
	GrandTotalAmount   int    `json:"grand_total_amount"`   // 优惠总金额（分）
	DecreaseAmount     int    `json:"decrease_amount"`      // 优惠减免金额（分）
	TransferAmount     int    `json:"transfer_amount"`      // 可转赠金额（分）
	CouponName         string `json:"coupon_name"`          // 代金券名称
	AvailableBeginTime string `json:"available_begin_time"` // 生效时间
	AvailableEndTime   string `json:"available_end_time"`   // 失效时间
	ReceiveTime        string `json:"receive_time"`         // 领取时间
	ReceiveMethod      string `json:"receive_method"`       // 领取方式
	SendSource         string `json:"send_source"`          // 发放来源
}
