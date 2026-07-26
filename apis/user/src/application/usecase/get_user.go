package usecase

import (
	"context"
	"errors"

	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/user/src/application/output-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	userrepo "github.com/sukuname4976/portfolio/apis/user/src/domain/repository-interfaces/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
)

// GetUserUseCase ユーザー取得ユースケースのインターフェース
type GetUserUseCase interface {
	Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.GetUserOutput, error)
}

type getUserUseCase struct {
	repo userrepo.Repository
}

// NewGetUserUseCase GetUserUseCaseを生成
func NewGetUserUseCase(repo userrepo.Repository) GetUserUseCase {
	return &getUserUseCase{repo: repo}
}

// Execute ユーザー取得を実行（DBから参照）
func (u *getUserUseCase) Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.GetUserOutput, error) {
	id, err := userid.New(input.ID)
	if err != nil {
		return nil, apperrors.NewUserNotFoundError("invalid user id: " + input.ID)
	}

	found, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, userrepo.ErrNotFound) {
			return nil, apperrors.NewUserNotFoundError("user not found: " + input.ID)
		}
		return nil, err
	}

	return u.toOutput(found), nil
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
