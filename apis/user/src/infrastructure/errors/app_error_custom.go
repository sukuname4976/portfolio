package errors

import "net/http"

// UserNotFoundError ユーザーが見つからないエラー
type UserNotFoundError struct {
	NotFoundError
}

// NewUserNotFoundError ユーザー未発見エラーを生成
func NewUserNotFoundError(internalMsg string) *UserNotFoundError {
	return &UserNotFoundError{
		NotFoundError: NotFoundError{
			BaseError: BaseError{
				ErrorID:              ErrIDUserNotFound,
				Status:               http.StatusNotFound,
				ExternalErrorMessage: "user not found",
				InternalErrorMessage: internalMsg,
			},
		},
	}
}

// UserValidationError ユーザーバリデーションエラー
type UserValidationError struct {
	BadRequestError
}

// NewUserValidationError ユーザーバリデーションエラーを生成
func NewUserValidationError(externalMsg, internalMsg string) *UserValidationError {
	return &UserValidationError{
		BadRequestError: BadRequestError{
			BaseError: BaseError{
				ErrorID:              ErrIDUserValidation,
				Status:               http.StatusBadRequest,
				ExternalErrorMessage: externalMsg,
				InternalErrorMessage: internalMsg,
			},
		},
	}
}
