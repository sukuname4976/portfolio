package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// responseWriter ステータスコードを記録するラッパー
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging リクエストログを出力するミドルウェア
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// ステータスコード記録用ラッパー
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(rw, r)

		// OpenTelemetryのtrace_idを取得（設定されていれば）
		var traceID string
		if span := trace.SpanFromContext(r.Context()); span.SpanContext().IsValid() {
			traceID = span.SpanContext().TraceID().String()
		}

		// 構造化ログ出力
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.status,
			"duration_ms", time.Since(start).Milliseconds(),
			"trace_id", traceID,
		)
	})
}
