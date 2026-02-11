package user

import (
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
)

// User ユーザーエンティティ
type User struct {
	id    userid.UserID
	name  string
	email email.Email
}

// New Userを生成
func New(id userid.UserID, name string, email email.Email) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

// ID ユーザーIDを取得
func (u *User) ID() userid.UserID {
	return u.id
}

// Name 名前を取得
func (u *User) Name() string {
	return u.name
}

// Email メールアドレスを取得
func (u *User) Email() email.Email {
	return u.email
}
