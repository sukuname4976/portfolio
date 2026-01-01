package main

import (
	"log"
	"net/http"

	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/config"
	"github.com/sukuname4976/portfolio/apis/bff/presentation/middleware"
	"github.com/sukuname4976/portfolio/apis/bff/presentation/router"
)

func main() {
	// 1. 設定の読み込み
	cfg := config.Load()

	// 2. ルーティング設定
	mux := http.NewServeMux()
	router.Setup(mux, cfg)

	// 3. ミドルウェアチェーンの構築
	handler := middleware.Recovery(
		middleware.Logging(
			middleware.JSON(mux),
		),
	)

	// 4. サーバー起動
	log.Printf("BFF Service starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}
