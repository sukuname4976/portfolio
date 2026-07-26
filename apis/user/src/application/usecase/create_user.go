package usecase

import (
	"context"
	"errors"

	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/user/src/application/output-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	userrepo "github.com/sukuname4976/portfolio/apis/user/src/domain/repository-interfaces/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
)

// CreateUserUseCase ユーザー登録ユースケースのインターフェース
type CreateUserUseCase interface {
	Execute(ctx context.Context, input inputdto.CreateUserInput) (*outputdto.CreateUserOutput, error)
}

type createUserUseCase struct {
	repo userrepo.Repository
}

// NewCreateUserUseCase CreateUserUseCaseを生成
func NewCreateUserUseCase(repo userrepo.Repository) CreateUserUseCase {
	return &createUserUseCase{repo: repo}
}

// Execute ユーザー登録を実行（DBへINSERT）
func (u *createUserUseCase) Execute(ctx context.Context, input inputdto.CreateUserInput) (*outputdto.CreateUserOutput, error) {
	if input.Name == "" {
		return nil, apperrors.NewUserValidationError("name is required", "empty name")
	}
	mail, err := email.New(input.Email)
	if err != nil {
		return nil, apperrors.NewUserValidationError("invalid email", err.Error())
	}

	created, err := u.repo.Create(ctx, input.Name, mail)
	if err != nil {
		if errors.Is(err, userrepo.ErrDuplicateEmail) {
			return nil, apperrors.NewUserConflictError("email already exists", err.Error())
		}
		return nil, err
	}

	return u.toOutput(created), nil
}

func (u *createUserUseCase) toOutput(user *user.User) *outputdto.CreateUserOutput {
	return &outputdto.CreateUserOutput{
		User: outputdto.UserDTO{
			ID:    user.ID().Value(),
			Name:  user.Name(),
			Email: user.Email().Value(),
		},
	}
}
