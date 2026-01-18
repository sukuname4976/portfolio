package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// RunEnv 実行環境
type RunEnv string

const (
	RunEnvLocal       RunEnv = "local"
	RunEnvTest        RunEnv = "test"
	RunEnvDevelopment RunEnv = "development"
	RunEnvProduction  RunEnv = "production"
)

// ParseRunEnv 文字列からRunEnvをパース
func ParseRunEnv(s string) (RunEnv, error) {
	switch RunEnv(s) {
	case RunEnvLocal:
		return RunEnvLocal, nil
	case RunEnvTest:
		return RunEnvTest, nil
	case RunEnvDevelopment:
		return RunEnvDevelopment, nil
	case RunEnvProduction:
		return RunEnvProduction, nil
	case "":
		return RunEnvLocal, nil
	default:
		return "", fmt.Errorf("invalid RUN_ENV: %q", s)
	}
}

// Config アプリケーション設定
type Config struct {
	RunEnv         RunEnv
	Port           string
	PokeAPIBaseURL string
	HTTPTimeout    time.Duration
}

// Load 実行環境に応じて設定を読み込み
func Load() (*Config, error) {
	runEnv, err := ParseRunEnv(os.Getenv("RUN_ENV"))
	if err != nil {
		return nil, err
	}

	// 環境に応じた設定ソースから読み込み
	switch runEnv {
	case RunEnvLocal, RunEnvTest:
		// ローカル/テスト環境: env/.envファイルから読み込み
		_ = godotenv.Load("env/.env")
	case RunEnvDevelopment, RunEnvProduction:
		// 開発/本番環境: Secret Managerから読み込み（未実装）
		// TODO: Google Cloud Secret Managerから読み込み
	}

	return &Config{
		RunEnv:         runEnv,
		Port:           getEnvOrDefault("PORT", "8000"),
		PokeAPIBaseURL: getEnvOrDefault("POKEAPI_BASE_URL", "https://pokeapi.co/api/v2"),
		HTTPTimeout:    10 * time.Second,
	}, nil
}

// getEnvOrDefault 環境変数を取得、なければデフォルト値を返す
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
