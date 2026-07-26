package user

// User ユーザー情報を表すエンティティ（user サービスから取得した情報の保持用）
type User struct {
	id    string
	name  string
	email string
}

// New Userエンティティを生成
func New(id, name, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

// ID ユーザーIDを取得
func (u *User) ID() string {
	return u.id
}

// Name 名前を取得
func (u *User) Name() string {
	return u.name
}

// Email メールアドレスを取得
func (u *User) Email() string {
	return u.email
}
