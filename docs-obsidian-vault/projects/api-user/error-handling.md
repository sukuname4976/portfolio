# エラーハンドリング方針

BFF API におけるエラーハンドリングの設計方針。

## 設計原則

- 内部エラー詳細はクライアントに露出しない
- すべてのエラーは構造化ログに記録する
- エラーの種類に応じて適切なHTTP ステータスを返す

## エラーの発生源と分類

- HTTP層: プロトコルエラー (JSON 構文エラー、型不一致)
- Controller層: 入力バリデーション (空文字、フォーマット不正)
- UseCase層: ビジネスロジック (権限不足、状態不整合)
- Gateway層: 外部通信 (API 接続失敗、タイムアウト)
- 全レイヤー: システム障害 (panic)

## ファイル構成

### エラー定義 (infrastructure/errors/)

```text
infrastructure/errors/
├── app_error_base.go       # BaseError + HTTP status別の基礎型
├── app_error_custom.go     # ビジネス固有のカスタムエラー
├── error_id_base.go        # 基礎エラーID定数
└── error_id_custom.go      # カスタムエラーID定数
```

### ミドルウェア (presentation/middleware/)

```text
presentation/middleware/
├── log_errors.go               # エラーログ出力 (LogError関数)
├── log_requests.go             # リクエスト/レスポンスログ
├── recover_panic.go            # panic回復 → 500
└── handle_decode_failures.go   # JSONデコード失敗 → 400
```

## 処理フロー

```text
リクエスト
    │
    ▼
┌─────────────────────┐
│ recover_panic.go    │──panic──→ 500 + ログ
└─────────────────────┘
    │
    ▼
┌─────────────────────┐
│ log_requests.go     │  リクエスト/レスポンスログ
└─────────────────────┘
    │
    ▼
┌─────────────────────┐
│ ogen Server         │
│ + handle_decode_    │──デコード失敗──→ 400 + ログ
│   failures.go       │                  (内部詳細は隠蔽)
└─────────────────────┘
    │
    ▼
┌─────────────────────┐
│ Controller          │
│ + log_errors.go     │──バリデーション失敗──→ AppError → 400
│                     │──Gateway失敗────────→ AppError → 502
└─────────────────────┘
    │
    ▼
正常レスポンス
```

## AppError 構造

```go
type BaseError struct {
    ErrorID              string  // エラー識別子 (例: ERR_ECHO_VALIDATION)
    Status               int     // HTTP ステータスコード
    ExternalErrorMessage string  // クライアントに返すメッセージ
    InternalErrorMessage string  // ログに記録する内部詳細
}
```

## 使用例

### Controller内でのエラー処理

```go
// バリデーションエラー
appErr := apperrors.NewEchoValidationError(
    "message is required",     // ExternalMessage (クライアントへ)
    "empty message received",  // InternalMessage (ログへ)
)
middleware.LogError(appErr)
return &ogen.EchoBadRequest{Error: appErr.ExternalErrorMessage}, nil
```

### ログ出力形式

```json
{
  "level": "ERROR",
  "msg": "request failed",
  "timestamp": "2024-01-17T12:00:00Z",
  "error_id": "ERR_ECHO_VALIDATION",
  "status": 400,
  "external_message": "message is required",
  "internal_message": "empty message received"
}
```

## エラー型の設計

基礎エラーとカスタムエラーの2層構造で設計する。

- 基礎エラー: HTTP ステータスコードに対応した汎用型 (BadRequestError, BadGatewayError など)
- カスタムエラー: 基礎エラーを埋め込み、ビジネス固有の意味を持たせた型

```go
// 基礎エラー (直接使わない)
type BadRequestError struct {
    BaseError
}

// カスタムエラー (実際に使う)
type EchoValidationError struct {
    BadRequestError  // 基礎エラーを埋め込み
}
```

カスタムエラーを通じて基礎エラーの機能を継承し、エラーIDやメッセージをビジネス要件に合わせてカスタマイズする。
