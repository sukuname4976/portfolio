package usecase_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/bff/src/application/usecase"
	userentity "github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/user"
	usergateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/user"
	mocks "github.com/sukuname4976/portfolio/apis/bff/test/mocks/user"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	t.Run("正常系：gateway が返した登録ユーザーを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		created := userentity.New("11111111-1111-1111-1111-111111111111", "田中 太郎", "tanaka@example.com")
		mockGW := mocks.NewMockGateway(ctrl)
		mockGW.EXPECT().
			Create(gomock.Any(), "田中 太郎", "tanaka@example.com").
			Return(created, nil)

		uc := usecase.NewCreateUserUseCase(mockGW)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "田中 太郎",
			Email: "tanaka@example.com",
		})

		require.NoError(t, err)
		assert.Equal(t, "11111111-1111-1111-1111-111111111111", output.ID)
		assert.Equal(t, "田中 太郎", output.Name)
		assert.Equal(t, "tanaka@example.com", output.Email)
	})

	t.Run("異常系：gateway の GatewayError をそのまま伝播する", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGW := mocks.NewMockGateway(ctrl)
		mockGW.EXPECT().
			Create(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, &usergateway.GatewayError{StatusCode: http.StatusConflict, Message: "email already exists"})

		uc := usecase.NewCreateUserUseCase(mockGW)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "田中 太郎",
			Email: "tanaka@example.com",
		})

		assert.Nil(t, output)
		var gatewayErr *usergateway.GatewayError
		require.ErrorAs(t, err, &gatewayErr)
		assert.Equal(t, http.StatusConflict, gatewayErr.StatusCode)
	})
}

func TestGetUserUseCase_Execute(t *testing.T) {
	t.Run("正常系：gateway が返したユーザーを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		found := userentity.New("11111111-1111-1111-1111-111111111111", "田中 太郎", "tanaka@example.com")
		mockGW := mocks.NewMockGateway(ctrl)
		mockGW.EXPECT().
			FetchByID(gomock.Any(), "11111111-1111-1111-1111-111111111111").
			Return(found, nil)

		uc := usecase.NewGetUserUseCase(mockGW)
		output, err := uc.Execute(context.Background(), inputdto.GetUserInput{ID: "11111111-1111-1111-1111-111111111111"})

		require.NoError(t, err)
		assert.Equal(t, "田中 太郎", output.Name)
		assert.Equal(t, "tanaka@example.com", output.Email)
	})

	t.Run("異常系：gateway の 404 GatewayError を伝播する", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGW := mocks.NewMockGateway(ctrl)
		mockGW.EXPECT().
			FetchByID(gomock.Any(), gomock.Any()).
			Return(nil, &usergateway.GatewayError{StatusCode: http.StatusNotFound, Message: "user not found"})

		uc := usecase.NewGetUserUseCase(mockGW)
		output, err := uc.Execute(context.Background(), inputdto.GetUserInput{ID: "missing"})

		assert.Nil(t, output)
		var gatewayErr *usergateway.GatewayError
		require.ErrorAs(t, err, &gatewayErr)
		assert.Equal(t, http.StatusNotFound, gatewayErr.StatusCode)
	})
}
