package controller

import (
	"context"
	"errors"

	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/application/usecase"
	userrepo "github.com/sukuname4976/portfolio/apis/user/src/domain/repository-interfaces/user"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
	ogen "github.com/sukuname4976/portfolio/apis/user/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/user/src/presentation/middleware"
)

// Handler ogenのHandlerインターフェースを実装
type Handler struct {
	getUserUseCase    usecase.GetUserUseCase
	createUserUseCase usecase.CreateUserUseCase
}

// NewHandler Handlerを生成（ユースケースは Repository から構築）
func NewHandler(repo userrepo.Repository) *Handler {
	return &Handler{
		getUserUseCase:    usecase.NewGetUserUseCase(repo),
		createUserUseCase: usecase.NewCreateUserUseCase(repo),
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

// CreateUser ユーザー登録エンドポイント
func (h *Handler) CreateUser(ctx context.Context, req *ogen.CreateUserRequest) (ogen.CreateUserRes, error) {
	input := inputdto.CreateUserInput{
		Name:  req.Name,
		Email: req.Email,
	}
	output, err := h.createUserUseCase.Execute(ctx, input)
	if err != nil {
		// バリデーションエラーは400
		var validationErr *apperrors.UserValidationError
		if errors.As(err, &validationErr) {
			middleware.LogError(validationErr)
			return &ogen.CreateUserBadRequest{
				Error: validationErr.GetExternalMessage(),
			}, nil
		}

		// email 重複は409
		var conflictErr *apperrors.UserConflictError
		if errors.As(err, &conflictErr) {
			middleware.LogError(conflictErr)
			return &ogen.CreateUserConflict{
				Error: conflictErr.GetExternalMessage(),
			}, nil
		}

		// その他のエラーは500
		appErr := apperrors.NewInternalServerError(err.Error())
		middleware.LogError(appErr)
		return &ogen.CreateUserInternalServerError{
			Error: appErr.GetExternalMessage(),
		}, nil
	}

	// 201 Created（UserResponse が createUser の 201 にマッピングされる）
	return &ogen.UserResponse{
		User: ogen.UserData{
			ID:    output.User.ID,
			Name:  output.User.Name,
			Email: output.User.Email,
		},
	}, nil
}
