package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-faster/jx"
)

// HandleDecodeFailures JSONデコード失敗時に安全なエラーレスポンスを返す
// 内部エラー詳細はログに記録し、クライアントには汎用メッセージのみ返す
func HandleDecodeFailures(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	// 構造化ログで内部詳細を記録
	slog.Error("decode failure",
		"timestamp", time.Now().UTC().Format(time.RFC3339),
		"error_id", "ERR_DECODE_FAILURE",
		"status", http.StatusBadRequest,
		"external_message", "invalid request format",
		"internal_message", err.Error(),
		"path", r.URL.Path,
		"method", r.Method,
	)

	// クライアントには汎用メッセージのみ返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	e := jx.GetEncoder()
	defer jx.PutEncoder(e)

	e.ObjStart()
	e.FieldStart("error")
	e.StrEscape("invalid request format")
	e.ObjEnd()

	_, _ = w.Write(e.Bytes())
}
