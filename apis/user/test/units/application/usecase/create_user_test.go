package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	inputdto "github.com/sukuname4976/portfolio/apis/user/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/user/src/application/usecase"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	userrepo "github.com/sukuname4976/portfolio/apis/user/src/domain/repository-interfaces/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
	apperrors "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/errors"
	mocks "github.com/sukuname4976/portfolio/apis/user/test/mocks/user"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	t.Run("正常系：登録したユーザーを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const id = "11111111-1111-1111-1111-111111111111"
		uid, _ := userid.New(id)
		mail, _ := email.New("tanaka@example.com")
		created := user.New(uid, "田中 太郎", mail)

		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().
			Create(gomock.Any(), "田中 太郎", gomock.Any()).
			Return(created, nil)

		uc := usecase.NewCreateUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "田中 太郎",
			Email: "tanaka@example.com",
		})

		require.NoError(t, err)
		assert.Equal(t, id, output.User.ID)
		assert.Equal(t, "田中 太郎", output.User.Name)
		assert.Equal(t, "tanaka@example.com", output.User.Email)
	})

	t.Run("異常系：名前が空だと UserValidationError (リポジトリは呼ばれない)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockRepository(ctrl)

		uc := usecase.NewCreateUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "",
			Email: "tanaka@example.com",
		})

		assert.Nil(t, output)
		var validationErr *apperrors.UserValidationError
		assert.ErrorAs(t, err, &validationErr)
	})

	t.Run("異常系：メールが不正だと UserValidationError (リポジトリは呼ばれない)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockRepository(ctrl)

		uc := usecase.NewCreateUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "田中 太郎",
			Email: "invalid",
		})

		assert.Nil(t, output)
		var validationErr *apperrors.UserValidationError
		assert.ErrorAs(t, err, &validationErr)
	})

	t.Run("異常系：email 重複だと UserConflictError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().
			Create(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, userrepo.ErrDuplicateEmail)

		uc := usecase.NewCreateUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.CreateUserInput{
			Name:  "田中 太郎",
			Email: "tanaka@example.com",
		})

		assert.Nil(t, output)
		var conflictErr *apperrors.UserConflictError
		assert.ErrorAs(t, err, &conflictErr)
	})
}
