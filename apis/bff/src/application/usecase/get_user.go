package usecase

import (
	"context"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/output-dto"
	usergateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/user"
)

// コンパイル時にインターフェース実装を検証
var _ GetUserUseCase = (*getUserUseCase)(nil)

// GetUserUseCase ユーザー取得ユースケースインターフェース
type GetUserUseCase interface {
	Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.UserOutput, error)
}

// getUserUseCase GetUserUseCaseの実装
type getUserUseCase struct {
	userGateway usergateway.Gateway
}

// NewGetUserUseCase GetUserUseCaseを生成
func NewGetUserUseCase(userGateway usergateway.Gateway) GetUserUseCase {
	return &getUserUseCase{userGateway: userGateway}
}

// Execute user サービスへ取得を中継する
func (u *getUserUseCase) Execute(ctx context.Context, input inputdto.GetUserInput) (*outputdto.UserOutput, error) {
	found, err := u.userGateway.FetchByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return toUserOutput(found), nil
}
