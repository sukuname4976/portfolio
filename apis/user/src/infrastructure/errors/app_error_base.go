package errors

import "net/http"

// BaseError 全エラーの基底型
type BaseError struct {
	ErrorID              string
	Status               int
	ExternalErrorMessage string
	InternalErrorMessage string
}

// Error errorインターフェースを実装
func (e *BaseError) Error() string {
	return e.InternalErrorMessage
}

// GetErrorID エラーIDを取得
func (e *BaseError) GetErrorID() string {
	return e.ErrorID
}

// GetStatus HTTPステータスコードを取得
func (e *BaseError) GetStatus() int {
	return e.Status
}

// GetExternalMessage 外部向けメッセージを取得
func (e *BaseError) GetExternalMessage() string {
	return e.ExternalErrorMessage
}

// GetInternalMessage 内部ログ向けメッセージを取得
func (e *BaseError) GetInternalMessage() string {
	return e.InternalErrorMessage
}

// =============================================================================
// 4xx クライアントエラー
// =============================================================================

// BadRequestError 400 Bad Request
type BadRequestError struct {
	BaseError
}

// NewBadRequestError 400エラーを生成
func NewBadRequestError(externalMsg, internalMsg string) *BadRequestError {
	return &BadRequestError{
		BaseError: BaseError{
			ErrorID:              ErrIDBadRequest,
			Status:               http.StatusBadRequest,
			ExternalErrorMessage: externalMsg,
			InternalErrorMessage: internalMsg,
		},
	}
}

// UnauthorizedError 401 Unauthorized
type UnauthorizedError struct {
	BaseError
}

// NewUnauthorizedError 401エラーを生成
func NewUnauthorizedError(internalMsg string) *UnauthorizedError {
	return &UnauthorizedError{
		BaseError: BaseError{
			ErrorID:              ErrIDUnauthorized,
			Status:               http.StatusUnauthorized,
			ExternalErrorMessage: "unauthorized",
			InternalErrorMessage: internalMsg,
		},
	}
}

// ForbiddenError 403 Forbidden
type ForbiddenError struct {
	BaseError
}

// NewForbiddenError 403エラーを生成
func NewForbiddenError(internalMsg string) *ForbiddenError {
	return &ForbiddenError{
		BaseError: BaseError{
			ErrorID:              ErrIDForbidden,
			Status:               http.StatusForbidden,
			ExternalErrorMessage: "forbidden",
			InternalErrorMessage: internalMsg,
		},
	}
}

// NotFoundError 404 Not Found
type NotFoundError struct {
	BaseError
}

// NewNotFoundError 404エラーを生成
func NewNotFoundError(externalMsg, internalMsg string) *NotFoundError {
	return &NotFoundError{
		BaseError: BaseError{
			ErrorID:              ErrIDNotFound,
			Status:               http.StatusNotFound,
			ExternalErrorMessage: externalMsg,
			InternalErrorMessage: internalMsg,
		},
	}
}

// ConflictError 409 Conflict
type ConflictError struct {
	BaseError
}

// NewConflictError 409エラーを生成
func NewConflictError(externalMsg, internalMsg string) *ConflictError {
	return &ConflictError{
		BaseError: BaseError{
			ErrorID:              ErrIDConflict,
			Status:               http.StatusConflict,
			ExternalErrorMessage: externalMsg,
			InternalErrorMessage: internalMsg,
		},
	}
}

// =============================================================================
// 5xx サーバーエラー
// =============================================================================

// InternalServerError 500 Internal Server Error
type InternalServerError struct {
	BaseError
}

// NewInternalServerError 500エラーを生成
func NewInternalServerError(internalMsg string) *InternalServerError {
	return &InternalServerError{
		BaseError: BaseError{
			ErrorID:              ErrIDInternalServer,
			Status:               http.StatusInternalServerError,
			ExternalErrorMessage: "internal server error",
			InternalErrorMessage: internalMsg,
		},
	}
}

// BadGatewayError 502 Bad Gateway
type BadGatewayError struct {
	BaseError
}

// NewBadGatewayError 502エラーを生成
func NewBadGatewayError(externalMsg, internalMsg string) *BadGatewayError {
	return &BadGatewayError{
		BaseError: BaseError{
			ErrorID:              ErrIDBadGateway,
			Status:               http.StatusBadGateway,
			ExternalErrorMessage: externalMsg,
			InternalErrorMessage: internalMsg,
		},
	}
}

// ServiceUnavailableError 503 Service Unavailable
type ServiceUnavailableError struct {
	BaseError
}

// NewServiceUnavailableError 503エラーを生成
func NewServiceUnavailableError(internalMsg string) *ServiceUnavailableError {
	return &ServiceUnavailableError{
		BaseError: BaseError{
			ErrorID:              ErrIDServiceUnavailable,
			Status:               http.StatusServiceUnavailable,
			ExternalErrorMessage: "service unavailable",
			InternalErrorMessage: internalMsg,
		},
	}
}

// GatewayTimeoutError 504 Gateway Timeout
type GatewayTimeoutError struct {
	BaseError
}

// NewGatewayTimeoutError 504エラーを生成
func NewGatewayTimeoutError(internalMsg string) *GatewayTimeoutError {
	return &GatewayTimeoutError{
		BaseError: BaseError{
			ErrorID:              ErrIDGatewayTimeout,
			Status:               http.StatusGatewayTimeout,
			ExternalErrorMessage: "gateway timeout",
			InternalErrorMessage: internalMsg,
		},
	}
}
