package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"

	"go.opentelemetry.io/otel/trace"
)

// Recovery パニックをキャッチしてエラーレスポンスを返すミドルウェア
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// OpenTelemetryのtrace_idを取得
				var traceID string
				if span := trace.SpanFromContext(r.Context()); span.SpanContext().IsValid() {
					traceID = span.SpanContext().TraceID().String()
				}

				// 構造化エラーログ
				slog.Error("panic recovered",
					"error", err,
					"method", r.Method,
					"path", r.URL.Path,
					"trace_id", traceID,
					"stack", string(debug.Stack()),
				)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				response := map[string]string{"error": "internal server error"}
				if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
					slog.Error("failed to encode error response", "error", encodeErr)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
