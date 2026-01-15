package errors

import "fmt"

// APIError 统一 API 错误类型
type APIError struct {
	StatusCode int    // HTTP 状态码
	Code       string // 业务错误码
	Message    string // 错误消息
	RequestID  string // 请求 ID
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("API error (status=%d, code=%s, request_id=%s): %s",
			e.StatusCode, e.Code, e.RequestID, e.Message)
	}
	return fmt.Sprintf("API error (status=%d, code=%s): %s",
		e.StatusCode, e.Code, e.Message)
}

// Is 实现 errors.Is 接口，支持错误比较
func (e *APIError) Is(target error) bool {
	t, ok := target.(*APIError)
	if !ok {
		return false
	}
	return e.Code == t.Code && e.StatusCode == t.StatusCode
}
