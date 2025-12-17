package api

import (
	"fmt"
	"time"
)

// WiFiAPI WiFi 管理 API
type WiFiAPI struct {
	*BaseAPI
}

// NewWiFiAPI 创建 WiFi 管理 API
func NewWiFiAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *WiFiAPI {
	return &WiFiAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// wifiToDateStr 将日期转换为字符串
func wifiToDateStr(date interface{}) (string, error) {
	switch v := date.(type) {
	case time.Time:
		return v.Format("2006-01-02"), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("cannot convert %T type to date string", v)
	}
}

// ListShops 获取门店列表
// http://mp.weixin.qq.com/wiki/15/bcfb5d4578ea818b89913472cf2bbf8f.html
func (api *WiFiAPI) ListShops(pageIndex, pageSize int) ([]map[string]interface{}, error) {
	result, err := api.Post("/bizwifi/shop/list", map[string]interface{}{
		"pageindex": pageIndex,
		"pagesize":  pageSize,
	})
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].([]interface{}); ok {
		shops := make([]map[string]interface{}, len(data))
		for i, item := range data {
			if shop, ok := item.(map[string]interface{}); ok {
				shops[i] = shop
			}
		}
		return shops, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetShop 查询门店的 WiFi 信息
// http://mp.weixin.qq.com/wiki/15/bcfb5d4578ea818b89913472cf2bbf8f.html
func (api *WiFiAPI) GetShop(shopID int) (map[string]interface{}, error) {
	result, err := api.Post("/bizwifi/shop/get", map[string]interface{}{
		"shop_id": shopID,
	})
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		return data, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// AddDevice 添加设备
// http://mp.weixin.qq.com/wiki/10/6232005bdc497f7cf8e19d4e843c70d2.html
func (api *WiFiAPI) AddDevice(shopID int, ssid, password, bssid string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"shop_id":  shopID,
		"ssid":     ssid,
		"password": password,
		"bssid":    bssid,
	}
	return api.Post("/bizwifi/device/add", data)
}

// ListDevices 查询设备
// http://mp.weixin.qq.com/wiki/10/6232005bdc497f7cf8e19d4e843c70d2.html
func (api *WiFiAPI) ListDevices(shopID *int, pageIndex, pageSize int) ([]map[string]interface{}, error) {
	data := make(map[string]interface{})
	if shopID != nil {
		data["shop_id"] = *shopID
	}
	data["pageindex"] = pageIndex
	data["pagesize"] = pageSize

	result, err := api.Post("/bizwifi/device/list", data)
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].([]interface{}); ok {
		devices := make([]map[string]interface{}, len(data))
		for i, item := range data {
			if device, ok := item.(map[string]interface{}); ok {
				devices[i] = device
			}
		}
		return devices, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// DeleteDevice 删除设备
// http://mp.weixin.qq.com/wiki/10/6232005bdc497f7cf8e19d4e843c70d2.html
func (api *WiFiAPI) DeleteDevice(bssid string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"bssid": bssid,
	}
	return api.Post("/bizwifi/device/delete", data)
}

// GetQRCodeURL 获取物料二维码图片网址
// http://mp.weixin.qq.com/wiki/7/fcd0378ef00617fc276be2b3baa80973.html
func (api *WiFiAPI) GetQRCodeURL(shopID, imgID int) (string, error) {
	result, err := api.Post("/bizwifi/qrcode/get", map[string]interface{}{
		"shop_id": shopID,
		"img_id":  imgID,
	})
	if err != nil {
		return "", err
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if qrcodeURL, ok := data["qrcode_url"].(string); ok {
			return qrcodeURL, nil
		}
	}
	return "", fmt.Errorf("unexpected response format")
}

// SetHomepage 设置商家主页
// http://mp.weixin.qq.com/wiki/6/2732f3cf83947e0e4971aa8797ee9d6a.html
func (api *WiFiAPI) SetHomepage(shopID, templateID int, url *string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"shop_id":     shopID,
		"template_id": templateID,
	}
	if url != nil {
		data["struct"] = map[string]interface{}{
			"url": *url,
		}
	}
	return api.Post("/bizwifi/homepage/set", data)
}

// GetHomepage 查询商家主页
// http://mp.weixin.qq.com/wiki/6/2732f3cf83947e0e4971aa8797ee9d6a.html
func (api *WiFiAPI) GetHomepage(shopID int) (map[string]interface{}, error) {
	result, err := api.Post("/bizwifi/homepage/get", map[string]interface{}{
		"shop_id": shopID,
	})
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		return data, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// ListStatistics Wi-Fi 数据统计
// http://mp.weixin.qq.com/wiki/8/dfa2b756b66fca5d9b1211bc18812698.html
func (api *WiFiAPI) ListStatistics(beginDate, endDate interface{}, shopID int) ([]map[string]interface{}, error) {
	beginDateStr, err := wifiToDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := wifiToDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/bizwifi/statistics/list", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
		"shop_id":    shopID,
	})
	if err != nil {
		return nil, err
	}

	if data, ok := result["data"].([]interface{}); ok {
		statistics := make([]map[string]interface{}, len(data))
		for i, item := range data {
			if stat, ok := item.(map[string]interface{}); ok {
				statistics[i] = stat
			}
		}
		return statistics, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}
