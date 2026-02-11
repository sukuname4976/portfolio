package userid

import "errors"

var ErrEmpty = errors.New("user id cannot be empty")

// UserID ユーザーIDを表す値オブジェクト
type UserID struct {
	value string
}

// New UserIDを生成
func New(value string) (UserID, error) {
	if value == "" {
		return UserID{}, ErrEmpty
	}
	return UserID{value: value}, nil
}

// Value 値を取得
func (id UserID) Value() string {
	return id.value
}
