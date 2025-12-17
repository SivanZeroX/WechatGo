package client

import (
	"encoding/json"
	"fmt"
)

// OAAPI 办公应用API
type OAAPI struct {
	BaseAPI interface {
		Post(url string, data interface{}) (map[string]interface{}, error)
		Get(url string, params map[string]string) (map[string]interface{}, error)
	}
}

// NewOAAPI 创建办公应用API
func NewOAAPI(client interface {
	Get(url string, params map[string]string) (map[string]interface{}, error)
	Post(url string, data interface{}) (map[string]interface{}, error)
	GetAccessToken() (string, error)
}) *OAAPI {
	return &OAAPI{
		BaseAPI: client,
	}
}

// GetApprovalInfo 获取审批信息
// https://developer.work.weixin.qq.com/document/path/91552
func (api *OAAPI) GetApprovalInfo(req *GetApprovalInfoRequest) (*GetApprovalInfoResponse, error) {
	result, err := api.BaseAPI.Post("/oa/getapprovalinfo", req)
	if err != nil {
		return nil, err
	}

	var resp GetApprovalInfoResponse
	if err := json.Unmarshal([]byte(result["data"].(string)), &resp.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// GetApprovalDetail 获取审批详情
// https://developer.work.weixin.qq.com/document/path/91552
func (api *OAAPI) GetApprovalDetail(spNo string) (*GetApprovalDetailResponse, error) {
	result, err := api.BaseAPI.Get("/oa/getapprovaldetail", map[string]string{
		"sp_no": spNo,
	})
	if err != nil {
		return nil, err
	}

	var resp GetApprovalDetailResponse
	if err := json.Unmarshal([]byte(result["sp_detail"].(string)), &resp.SpDetail); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &resp, nil
}

// ==================== 请求结构体 ====================

// GetApprovalInfoRequest 获取审批信息请求
type GetApprovalInfoRequest struct {
	StartTime string `json:"starttime"` // 开始时间
	EndTime   string `json:"endtime"`   // 结束时间
	SpStatus  int    `json:"sp_status"` // 审批状态
}

// ==================== 响应结构体 ====================

// GetApprovalInfoResponse 获取审批信息响应
type GetApprovalInfoResponse struct {
	Data []ApprovalInfo `json:"data"` // 审批信息列表
}

// GetApprovalDetailResponse 获取审批详情响应
type GetApprovalDetailResponse struct {
	SpDetail ApprovalDetail `json:"sp_detail"` // 审批详情
}

// ApprovalInfo 审批信息
type ApprovalInfo struct {
	SpNo      string `json:"sp_no"`      // 审批单号
	SpName    string `json:"sp_name"`    // 审批模板名称
	Applicant string `json:"applyer"`    // 申请人
	Dept      string `json:"department"` // 部门
	Status    int    `json:"sp_status"`  // 审批状态
	ApplyTime string `json:"apply_time"` // 申请时间
}

// ApprovalDetail 审批详情
type ApprovalDetail struct {
	SpNo      string         `json:"sp_no"`      // 审批单号
	SpName    string         `json:"sp_name"`    // 审批模板名称
	Applicant string         `json:"applyer"`    // 申请人
	Dept      string         `json:"department"` // 部门
	Status    int            `json:"sp_status"`  // 审批状态
	ApplyTime string         `json:"apply_time"` // 申请时间
	Details   []ApprovalNode `json:"sp_detail"`  // 审批节点
}

// ApprovalNode 审批节点
type ApprovalNode struct {
	Approver   string `json:"approver"`    // 审批人
	Action     int    `json:"action"`      // 审批动作
	ActionTime string `json:"action_time"` // 审批时间
}
