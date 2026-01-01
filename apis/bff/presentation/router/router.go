package router

import (
	"fmt"
	"net/http"

	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/config"
	"github.com/sukuname4976/portfolio/apis/bff/presentation/controller"
)

// Setup ルーティングを設定
func Setup(mux *http.ServeMux, cfg *config.Config) {
	// ヘルスチェック
	mux.HandleFunc("/", healthHandler(cfg))

	// POST /echo
	mux.HandleFunc("/echo", controller.NewEchoHandler(cfg))
}

// healthHandler ヘルスチェックハンドラー
func healthHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "BFF Service is running on port %s\n", cfg.Port)
	}
}
