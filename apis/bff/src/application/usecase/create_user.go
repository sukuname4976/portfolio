package usecase

import (
	"context"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/output-dto"
	userentity "github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/user"
	usergateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/user"
)

// コンパイル時にインターフェース実装を検証
var _ CreateUserUseCase = (*createUserUseCase)(nil)

// CreateUserUseCase ユーザー登録ユースケースインターフェース
type CreateUserUseCase interface {
	Execute(ctx context.Context, input inputdto.CreateUserInput) (*outputdto.UserOutput, error)
}

// createUserUseCase CreateUserUseCaseの実装
type createUserUseCase struct {
	userGateway usergateway.Gateway
}

// NewCreateUserUseCase CreateUserUseCaseを生成
func NewCreateUserUseCase(userGateway usergateway.Gateway) CreateUserUseCase {
	return &createUserUseCase{userGateway: userGateway}
}

// Execute user サービスへ登録を中継する
func (u *createUserUseCase) Execute(ctx context.Context, input inputdto.CreateUserInput) (*outputdto.UserOutput, error) {
	created, err := u.userGateway.Create(ctx, input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	return toUserOutput(created), nil
}

// toUserOutput ドメインエンティティからOutputDTOに変換
func toUserOutput(u *userentity.User) *outputdto.UserOutput {
	return &outputdto.UserOutput{
		ID:    u.ID(),
		Name:  u.Name(),
		Email: u.Email(),
	}
}
