package user

import (
	"context"
	"errors"

	"github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
)

// infrastructure 層の DB ドライバ固有エラーを domain 非依存の形に変換するためのセンチネルエラー。
var (
	// ErrNotFound 指定ユーザーが存在しない。
	ErrNotFound = errors.New("user not found")
	// ErrDuplicateEmail email の UNIQUE 制約に違反した (既に同一 email が存在する)。
	ErrDuplicateEmail = errors.New("email already exists")
)

// Repository ユーザー永続化インターフェース
type Repository interface {
	// Create 新規ユーザーを永続化し、DB が発行した ID を含むユーザーを返す。
	Create(ctx context.Context, name string, email email.Email) (*user.User, error)
	// FindByID ID でユーザーを取得する。存在しない場合は ErrNotFound を返す。
	FindByID(ctx context.Context, id userid.UserID) (*user.User, error)
}
