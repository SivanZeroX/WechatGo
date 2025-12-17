package api

import (
	"fmt"
	"time"
)

// DataCubeAPI 数据统计 API
type DataCubeAPI struct {
	*BaseAPI
}

// NewDataCubeAPI 创建数据统计 API
func NewDataCubeAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *DataCubeAPI {
	return &DataCubeAPI{
		BaseAPI: NewBaseAPI(client),
	}
}

// toDateStr 将日期转换为字符串
func toDateStr(date interface{}) (string, error) {
	switch v := date.(type) {
	case time.Time:
		return v.Format("2006-01-02"), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("cannot convert %T type to date string", v)
	}
}

// GetUserSummary 获取用户增减数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserSummary(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getusersummary", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUserCumulate 获取累计用户数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserCumulate(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getusercumulate", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetInterfaceSummary 获取接口分析数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (api *DataCubeAPI) GetInterfaceSummary(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getinterfacesummary", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetInterfaceSummaryHour 获取接口分析分时数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (api *DataCubeAPI) GetInterfaceSummaryHour(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getinterfacesummaryhour", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetArticleSummary 获取图文群发每日数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetArticleSummary(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getarticlesummary", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetArticleTotal 获取图文群发总数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetArticleTotal(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getarticletotal", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUserRead 获取图文统计数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserRead(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getuserread", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUserReadHour 获取图文分时统计数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserReadHour(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getuserreadhour", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUserShare 获取图文分享转发数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserShare(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getusershare", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUserShareHour 获取图文分享转发分时数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (api *DataCubeAPI) GetUserShareHour(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getusersharehour", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsg 获取消息发送概况数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsg(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsg", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgHour 获取消息发送分时数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsgHour(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsghour", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgWeek 获取消息发送周数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsgWeek(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsgweek", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgMonth 获取消息发送月数据
// http://mp.weixin.qq.com/wiki/12/32d42ad542f2e4fc8a8aa60e1bce9838.html
func (api *DataCubeAPI) GetUpstreamMsgMonth(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsgmonth", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgDist 获取消息发送分布数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsgDist(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsgdist", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgDistWeek 获取消息发送分布周数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsgDistWeek(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsgdistweek", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}

// GetUpstreamMsgDistMonth 获取消息发送分布月数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (api *DataCubeAPI) GetUpstreamMsgDistMonth(beginDate, endDate interface{}) ([]map[string]interface{}, error) {
	beginDateStr, err := toDateStr(beginDate)
	if err != nil {
		return nil, err
	}
	endDateStr, err := toDateStr(endDate)
	if err != nil {
		return nil, err
	}

	result, err := api.Post("/datacube/getupstreammsgdistmonth", map[string]interface{}{
		"begin_date": beginDateStr,
		"end_date":   endDateStr,
	})
	if err != nil {
		return nil, err
	}

	if list, ok := result["list"].([]interface{}); ok {
		listMap := make([]map[string]interface{}, len(list))
		for i, item := range list {
			if itemMap, ok := item.(map[string]interface{}); ok {
				listMap[i] = itemMap
			}
		}
		return listMap, nil
	}
	return nil, fmt.Errorf("unexpected response format")
}
