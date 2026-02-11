package usecase

import (
	"context"

	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/user/src/application/output-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
)

// GetUserUseCase ユーザー取得ユースケースのインターフェース
type GetUserUseCase interface {
	Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.GetUserOutput, error)
}

type getUserUseCase struct{}

// NewGetUserUseCase GetUserUseCaseを生成
func NewGetUserUseCase() GetUserUseCase {
	return &getUserUseCase{}
}

// Execute ユーザー取得を実行（ダミーデータを返却）
func (u *getUserUseCase) Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.GetUserOutput, error) {
	// ダミーユーザーID
	const dummyUserID = "dummy-user-1"

	if input.ID != dummyUserID {
		return nil, apperrors.NewUserNotFoundError("user not found: " + input.ID)
	}

	// ダミーユーザーを生成
	id, _ := userid.New(dummyUserID)
	mail, _ := email.New("tanaka@example.com")
	dummyUser := user.New(id, "田中 太郎", mail)

	return u.toOutput(dummyUser), nil
}

func (u *getUserUseCase) toOutput(user *user.User) *outputdto.GetUserOutput {
	return &outputdto.GetUserOutput{
		User: outputdto.UserDTO{
			ID:    user.ID().Value(),
			Name:  user.Name(),
			Email: user.Email().Value(),
		},
	}
}
