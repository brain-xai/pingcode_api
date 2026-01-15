package sdk

import (
	stderrors "errors"

	internalevents "github.com/brain-xai/pingcode_api/internal/errors"
)

// APIError 统一 API 错误类型（重新导出 internal.errors.APIError）
type APIError = internalevents.APIError

// 预定义错误
var (
	// ErrInvalidConfig 配置错误
	ErrInvalidConfig = stderrors.New("invalid configuration")

	// ErrUnauthorized 认证失败
	ErrUnauthorized = &APIError{StatusCode: 401, Code: "unauthorized", Message: "authentication failed"}

	// ErrNotFound 资源未找到
	ErrNotFound = &APIError{StatusCode: 404, Code: "not_found", Message: "resource not found"}
)
