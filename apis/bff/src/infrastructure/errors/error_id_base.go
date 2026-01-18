package errors

// =============================================================================
// 基礎エラーID定義
// =============================================================================

// 4xx クライアントエラーID
const (
	ErrIDBadRequest   = "ERR_BAD_REQUEST"  // 400
	ErrIDUnauthorized = "ERR_UNAUTHORIZED" // 401
	ErrIDForbidden    = "ERR_FORBIDDEN"    // 403
	ErrIDNotFound     = "ERR_NOT_FOUND"    // 404
	ErrIDConflict     = "ERR_CONFLICT"     // 409
)

// 5xx サーバーエラーID
const (
	ErrIDInternalServer     = "ERR_INTERNAL_SERVER"     // 500
	ErrIDBadGateway         = "ERR_BAD_GATEWAY"         // 502
	ErrIDServiceUnavailable = "ERR_SERVICE_UNAVAILABLE" // 503
	ErrIDGatewayTimeout     = "ERR_GATEWAY_TIMEOUT"     // 504
)
