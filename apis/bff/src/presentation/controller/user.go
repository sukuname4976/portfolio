package controller

import (
	"context"
	"errors"
	"net/http"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	usergateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/user"
	apperrors "github.com/sukuname4976/portfolio/apis/bff/src/infrastructure/errors"
	ogen "github.com/sukuname4976/portfolio/apis/bff/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/bff/src/presentation/middleware"
)

// CreateUser POST /api/v1/users を user サービスへ中継する
func (h *Handler) CreateUser(ctx context.Context, req *ogen.CreateUserRequest) (ogen.CreateUserRes, error) {
	input := inputdto.CreateUserInput{
		Name:  req.Name,
		Email: req.Email,
	}
	output, err := h.createUserUseCase.Execute(ctx, input)
	if err != nil {
		// user サービスが返した 4xx はそのままクライアントへ透過する
		var gatewayErr *usergateway.GatewayError
		if errors.As(err, &gatewayErr) {
			switch gatewayErr.StatusCode {
			case http.StatusBadRequest:
				return &ogen.CreateUserBadRequest{Error: gatewayErr.Message}, nil
			case http.StatusConflict:
				return &ogen.CreateUserConflict{Error: gatewayErr.Message}, nil
			}
		}

		// 通信失敗・上流 5xx などは 502 を返す
		appErr := apperrors.NewUserGatewayError(err.Error())
		middleware.LogError(appErr)
		return &ogen.CreateUserBadGateway{Error: appErr.ExternalErrorMessage}, nil
	}

	// 201 Created（UserResponse が createUser の 201 にマッピングされる）
	return &ogen.UserResponse{
		User: ogen.UserData{
			ID:    output.ID,
			Name:  output.Name,
			Email: output.Email,
		},
	}, nil
}

// GetUser GET /api/v1/users/{id} を user サービスへ中継する
func (h *Handler) GetUser(ctx context.Context, params ogen.GetUserParams) (ogen.GetUserRes, error) {
	input := inputdto.GetUserInput{ID: params.ID}
	output, err := h.getUserUseCase.Execute(ctx, input)
	if err != nil {
		// user サービスが 404 を返したら透過する
		var gatewayErr *usergateway.GatewayError
		if errors.As(err, &gatewayErr) && gatewayErr.StatusCode == http.StatusNotFound {
			return &ogen.GetUserNotFound{Error: gatewayErr.Message}, nil
		}

		appErr := apperrors.NewUserGatewayError(err.Error())
		middleware.LogError(appErr)
		return &ogen.GetUserBadGateway{Error: appErr.ExternalErrorMessage}, nil
	}

	return &ogen.UserResponse{
		User: ogen.UserData{
			ID:    output.ID,
			Name:  output.Name,
			Email: output.Email,
		},
	}, nil
}
