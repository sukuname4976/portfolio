package controller

import (
	"context"
	"strings"

	ogen "github.com/sukuname4976/portfolio/apis/bff/presentation/auto-generated-by-ogen"
)

// HealthCheck GET / を処理
func (h *Handler) HealthCheck(ctx context.Context) (ogen.HealthCheckOK, error) {
	return ogen.HealthCheckOK{
		Data: strings.NewReader("BFF Service is running\n"),
	}, nil
}
