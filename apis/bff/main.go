package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/sukuname4976/portfolio/apis/bff/src/infrastructure/config"
	ogen "github.com/sukuname4976/portfolio/apis/bff/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/bff/src/presentation/controller"
	"github.com/sukuname4976/portfolio/apis/bff/src/presentation/middleware"
)

func main() {
	// 1. slog初期設定（JSON形式）
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// 2. 設定の読み込み
	cfg := config.Load()

	// 3. ogenサーバーの構築
	handler := controller.NewHandler(cfg)
	server, err := ogen.NewServer(handler)
	if err != nil {
		slog.Error("failed to create server", "error", err)
		os.Exit(1)
	}

	// 4. ミドルウェアチェーンの構築
	httpHandler := middleware.Recovery(
		middleware.Logging(server),
	)

	// 5. サーバー起動
	slog.Info("server starting", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, httpHandler); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
