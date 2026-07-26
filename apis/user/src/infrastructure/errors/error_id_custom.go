package errors

// =============================================================================
// カスタムエラーID定義
// =============================================================================

// User エンドポイント用
const (
	ErrIDUserNotFound   = "ERR_USER_NOT_FOUND"
	ErrIDUserValidation = "ERR_USER_VALIDATION"
	ErrIDUserConflict   = "ERR_USER_CONFLICT"
)
