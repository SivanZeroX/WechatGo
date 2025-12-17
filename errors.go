package wechatgo

import (
	"fmt"
	"net/http"
)

// Error 微信错误基类
type Error struct {
	ErrCode int
	ErrMsg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", e.ErrCode, e.ErrMsg)
}

// NewError 创建新的微信错误
func NewError(errcode int, errmsg string) *Error {
	return &Error{
		ErrCode: errcode,
		ErrMsg:  errmsg,
	}
}

// ClientError 微信 API 客户端错误
type ClientError struct {
	ErrCode  int
	ErrMsg   string
	Request  *http.Request
	Response *http.Response
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", e.ErrCode, e.ErrMsg)
}

// NewClientError 创建新的客户端错误
func NewClientError(errcode int, errmsg string, req *http.Request, resp *http.Response) *ClientError {
	return &ClientError{
		ErrCode:  errcode,
		ErrMsg:   errmsg,
		Request:  req,
		Response: resp,
	}
}

// InvalidSignatureError 无效签名错误
type InvalidSignatureError struct {
	ErrCode int
	ErrMsg  string
}

func (e *InvalidSignatureError) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", e.ErrCode, e.ErrMsg)
}

// NewInvalidSignatureError 创建无效签名错误
func NewInvalidSignatureError() *InvalidSignatureError {
	return &InvalidSignatureError{
		ErrCode: -40001,
		ErrMsg:  "Invalid signature",
	}
}

// APILimitedError API 调用频率受限错误
type APILimitedError struct {
	ClientError
}

// NewAPILimitedError 创建 API 受限错误
func NewAPILimitedError(errcode int, errmsg string, req *http.Request, resp *http.Response) *APILimitedError {
	return &APILimitedError{
		ClientError: *NewClientError(errcode, errmsg, req, resp),
	}
}

// InvalidAppIDError 无效 AppID 错误
type InvalidAppIDError struct {
	ErrCode int
	ErrMsg  string
}

func (e *InvalidAppIDError) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", e.ErrCode, e.ErrMsg)
}

// NewInvalidAppIDError 创建无效 AppID 错误
func NewInvalidAppIDError() *InvalidAppIDError {
	return &InvalidAppIDError{
		ErrCode: -40005,
		ErrMsg:  "Invalid AppId",
	}
}

// InvalidMchIDError 无效商户 ID 错误
type InvalidMchIDError struct {
	ErrCode int
	ErrMsg  string
}

func (e *InvalidMchIDError) Error() string {
	return fmt.Sprintf("Error code: %d, message: %s", e.ErrCode, e.ErrMsg)
}

// NewInvalidMchIDError 创建无效商户 ID 错误
func NewInvalidMchIDError() *InvalidMchIDError {
	return &InvalidMchIDError{
		ErrCode: -40006,
		ErrMsg:  "Invalid MchId",
	}
}

// OAuthError 微信 OAuth 错误
type OAuthError struct {
	ClientError
}

// NewOAuthError 创建 OAuth 错误
func NewOAuthError(errcode int, errmsg string, req *http.Request, resp *http.Response) *OAuthError {
	return &OAuthError{
		ClientError: *NewClientError(errcode, errmsg, req, resp),
	}
}

// ComponentOAuthError 微信第三方平台 OAuth 错误
type ComponentOAuthError struct {
	ClientError
}

// NewComponentOAuthError 创建第三方平台 OAuth 错误
func NewComponentOAuthError(errcode int, errmsg string, req *http.Request, resp *http.Response) *ComponentOAuthError {
	return &ComponentOAuthError{
		ClientError: *NewClientError(errcode, errmsg, req, resp),
	}
}

// PayError 微信支付错误
type PayError struct {
	ClientError
	ReturnCode string
	ResultCode string
	ReturnMsg  string
}

func (e *PayError) Error() string {
	return fmt.Sprintf("Error code: %s, message: %s. Pay Error code: %d, message: %s",
		e.ReturnCode, e.ReturnMsg, e.ErrCode, e.ErrMsg)
}

// NewPayError 创建支付错误
func NewPayError(returnCode, resultCode, returnMsg string, errcode int, errmsg string,
	req *http.Request, resp *http.Response) *PayError {
	return &PayError{
		ClientError: *NewClientError(errcode, errmsg, req, resp),
		ReturnCode:  returnCode,
		ResultCode:  resultCode,
		ReturnMsg:   returnMsg,
	}
}

// PayV3Error 微信支付 V3 错误
type PayV3Error struct {
	ClientError
	Code    string
	Message string
}

func (e *PayV3Error) Error() string {
	return fmt.Sprintf("Error code: %s, message: %s", e.Code, e.Message)
}

// NewPayV3Error 创建支付 V3 错误
func NewPayV3Error(code, message string, req *http.Request, resp *http.Response) *PayV3Error {
	return &PayV3Error{
		ClientError: *NewClientError(0, message, req, resp),
		Code:        code,
		Message:     message,
	}
}
