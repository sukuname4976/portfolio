package requestmodel

import "errors"

// ErrMessageRequired メッセージが必須の場合のエラー
var ErrMessageRequired = errors.New("message is required")

// EchoRequest /echo リクエストボディ
type EchoRequest struct {
	Message string `json:"message"`
}

// Validate リクエストのバリデーション
func (r *EchoRequest) Validate() error {
	if r.Message == "" {
		return ErrMessageRequired
	}
	return nil
}
