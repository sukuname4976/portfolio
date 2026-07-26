package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/sukuname4976/portfolio/apis/user/src/infrastructure/config"
	"github.com/sukuname4976/portfolio/apis/user/src/infrastructure/db"
	userrepository "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/repository/user"
	ogen "github.com/sukuname4976/portfolio/apis/user/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/user/src/presentation/controller"
	"github.com/sukuname4976/portfolio/apis/user/src/presentation/middleware"
)

func main() {
	// 1. slog初期設定（JSON形式）
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// 2. 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// 3. DB 接続プールの構築
	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	// 4. 依存の組み立て（Repository → ogen Handler）
	repo := userrepository.NewRepository(pool)
	handler := controller.NewHandler(repo)
	server, err := ogen.NewServer(
		handler,
		ogen.WithErrorHandler(middleware.HandleDecodeFailures),
	)
	if err != nil {
		slog.Error("failed to create server", "error", err)
		os.Exit(1)
	}

	// 5. ミドルウェアチェーンの構築
	httpHandler := middleware.Recovery(
		middleware.Logging(server),
	)

	// 6. サーバー起動
	slog.Info("server starting", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, httpHandler); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
