# BFF API

## 概要

Golang 製の BFF (Backend for Frontend) サービス。クリーンアーキテクチャに基づいて構築。

## 技術スタック

- Go 1.25
- [ogen](https://github.com/ogen-go/ogen) : OpenAPIコード生成
- slog : 構造化ログ
- testify + gomock : テスト

## ディレクトリ構成

```text
apis/bff/
├── IF/                    # 外部サービスからの参照を受ける OpenAPI スキーマ
│   ├── openapi.yaml
│   ├── paths/
│   └── components/
├── src/
│   ├── application/       # Usecase·DTO
│   ├── domain/            # ドメインモデル·インターフェース
│   ├── infrastructure/    # 外部サービス
│   └── presentation/      # コントローラー·リクエストレスポンスモデル
├── test/                  # 単体テスト·統合テスト·モック
├── Dockerfile             # コンテナ定義
├── docker-compose.ci.yaml # CI 用 Compose
├── docker-compose.yaml    # 開発用 Compose
├── main.go                # API エントリーポイント
├── Makefile               # コマンド
└── project.json           # Nx プロジェクト設定
```

## 開発環境のセットアップ

1. VSCode で `.devcontainer/go` の devcontainer を開く
2. ターミナルで以下を実行

```bash
cd apis/bff

# 依存関係ダウンロード
go mod download

# 開発サーバー起動する場合 (ホットリロード対応)
make dev
```

## make コマンド

- `make dev` : 開発時に利用するサーバー起動 (ホットリロード対応)
- `make up` : 本番同等条件を試すサーバー起動コマンド
- `make gen-ogen` : OpenAPIからコード生成
- `make gen-mock` : モック生成
- `make check` : CI チェック (ビルド + テスト + リンティング確認 + フォーマット確認)
- `make prepare` : コミット前の品質担保処理
                   (ビルド + テスト + リンティング修正 + フォーマット修正 + gen-ogen + gen-mock)

## API仕様

OpenAPIスキーマは `IF/openapi.yaml` を参照。

### 提供エンドポイント

- `GET /health` : ヘルスチェック
- `POST /echo` : 開発進捗後に破棄する開発·連携確認用の仮エンドポイント (エコー + ランダムポケモン情報)
