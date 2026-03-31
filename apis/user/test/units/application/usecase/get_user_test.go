package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/application/usecase"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
)

func TestGetUserUseCase_Execute(t *testing.T) {
	t.Run("正常系：dummy-user-1でダミーユーザーを返す", func(t *testing.T) {
		uc := usecase.NewGetUserUseCase()
		input := inputdto.GetUserInput{ID: "dummy-user-1"}

		output, err := uc.Execute(context.Background(), input)

		require.NoError(t, err)
		assert.Equal(t, "dummy-user-1", output.User.ID)
		assert.Equal(t, "田中 太郎", output.User.Name)
		assert.Equal(t, "tanaka@example.com", output.User.Email)
	})

	t.Run("異常系：存在しないIDでUserNotFoundError", func(t *testing.T) {
		uc := usecase.NewGetUserUseCase()
		input := inputdto.GetUserInput{ID: "unknown-user"}

		output, err := uc.Execute(context.Background(), input)

		assert.Nil(t, output)
		assert.Error(t, err)

		var notFoundErr *apperrors.UserNotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})
}
