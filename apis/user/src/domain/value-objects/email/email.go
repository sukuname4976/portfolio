package email

import (
	"errors"
	"strings"
)

var (
	ErrEmpty         = errors.New("email cannot be empty")
	ErrInvalidFormat = errors.New("email format is invalid")
)

// Email メールアドレスを表す値オブジェクト
type Email struct {
	value string
}

// New Emailを生成
func New(value string) (Email, error) {
	if value == "" {
		return Email{}, ErrEmpty
	}
	if !strings.Contains(value, "@") {
		return Email{}, ErrInvalidFormat
	}
	return Email{value: value}, nil
}

// Value 値を取得
func (e Email) Value() string {
	return e.value
}
