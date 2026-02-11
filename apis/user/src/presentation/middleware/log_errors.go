package middleware

import (
	"log/slog"
	"time"
)

// AppError アプリケーションエラーの共通インターフェース
type AppError interface {
	error
	GetErrorID() string
	GetStatus() int
	GetExternalMessage() string
	GetInternalMessage() string
}

// LogError AppErrorを構造化ログ出力する
func LogError(err AppError) {
	slog.Error("request failed",
		"timestamp", time.Now().UTC().Format(time.RFC3339),
		"error_id", err.GetErrorID(),
		"status", err.GetStatus(),
		"external_message", err.GetExternalMessage(),
		"internal_message", err.GetInternalMessage(),
	)
}
