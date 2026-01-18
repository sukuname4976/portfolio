package message

import "errors"

// ErrEmpty メッセージが空の場合のエラー
var ErrEmpty = errors.New("message cannot be empty")

// Message メッセージ値オブジェクト
type Message struct {
	value string
}

// New Messageを生成（空文字チェック）
func New(value string) (Message, error) {
	if value == "" {
		return Message{}, ErrEmpty
	}
	return Message{value: value}, nil
}

// Value 値を取得
func (m Message) Value() string {
	return m.value
}
