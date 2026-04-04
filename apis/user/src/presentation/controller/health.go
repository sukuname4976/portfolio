package controller

import (
	"context"
	"strings"

	ogen "github.com/sukuname4976/portfolio/apis/user/src/presentation/auto-generated-by-ogen"
)

// HealthCheck ヘルスチェックエンドポイント
func (h *Handler) HealthCheck(ctx context.Context) (ogen.HealthCheckOK, error) {
	return ogen.HealthCheckOK{
		Data: strings.NewReader("User Service is running\n"),
	}, nil
}
