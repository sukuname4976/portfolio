package config

import (
	"os"
	"time"
)

// Config アプリケーション設定
type Config struct {
	Port           string
	PokeAPIBaseURL string
	HTTPTimeout    time.Duration
}

// Load 環境変数から設定を読み込み
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pokeAPIBaseURL := os.Getenv("POKEAPI_BASE_URL")
	if pokeAPIBaseURL == "" {
		pokeAPIBaseURL = "https://pokeapi.co/api/v2"
	}

	return &Config{
		Port:           port,
		PokeAPIBaseURL: pokeAPIBaseURL,
		HTTPTimeout:    10 * time.Second,
	}
}
