# BFF API

Go製のBFF（Backend for Frontend）サービス。

## 技術スタック

- Go 1.21
- 標準ライブラリ（net/http）

## コーディングルール

- `go fmt` でフォーマット
- `go vet` でチェック
- エラーは必ずハンドリングする（`_`で握りつぶさない）
- 関数コメントは `// FunctionName ...` の形式で書く

## ディレクトリ構成

- `main.go` - エントリーポイント
- `IF/` - インターフェース定義（OpenAPI）

