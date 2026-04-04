package controller

import (
	"context"
	"errors"

	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/application/usecase"
	"github.com/sukuname4976/portfolio/apis/user/src/infrastructure/config"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
	ogen "github.com/sukuname4976/portfolio/apis/user/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/user/src/presentation/middleware"
)

// Handler ogenのHandlerインターフェースを実装
type Handler struct {
	getUserUseCase usecase.GetUserUseCase
}

// NewHandler Handlerを生成
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		getUserUseCase: usecase.NewGetUserUseCase(),
	}
}

// GetUser ユーザー取得エンドポイント
func (h *Handler) GetUser(ctx context.Context, params ogen.GetUserParams) (ogen.GetUserRes, error) {
	input := inputdto.GetUserInput{ID: params.ID}
	output, err := h.getUserUseCase.Execute(ctx, input)
	if err != nil {
		// UserNotFoundErrorの場合は404
		var notFoundErr *apperrors.UserNotFoundError
		if errors.As(err, &notFoundErr) {
			middleware.LogError(notFoundErr)
			return &ogen.GetUserNotFound{
				Error: notFoundErr.GetExternalMessage(),
			}, nil
		}

		// その他のエラーは500
		appErr := apperrors.NewInternalServerError(err.Error())
		middleware.LogError(appErr)
		return &ogen.GetUserInternalServerError{
			Error: appErr.GetExternalMessage(),
		}, nil
	}

	return &ogen.UserResponse{
		User: ogen.UserData{
			ID:    output.User.ID,
			Name:  output.User.Name,
			Email: output.User.Email,
		},
	}, nil
}
