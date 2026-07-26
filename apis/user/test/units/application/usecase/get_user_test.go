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

func TestGetUserUseCase_Execute(t *testing.T) {
	t.Run("正常系：リポジトリが返したユーザーを返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		const id = "11111111-1111-1111-1111-111111111111"
		uid, _ := userid.New(id)
		mail, _ := email.New("tanaka@example.com")
		found := user.New(uid, "田中 太郎", mail)

		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().
			FindByID(gomock.Any(), gomock.Any()).
			Return(found, nil)

		uc := usecase.NewGetUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.GetUserInput{ID: id})

		require.NoError(t, err)
		assert.Equal(t, id, output.User.ID)
		assert.Equal(t, "田中 太郎", output.User.Name)
		assert.Equal(t, "tanaka@example.com", output.User.Email)
	})

	t.Run("異常系：リポジトリが ErrNotFound を返すと UserNotFoundError", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockRepository(ctrl)
		mockRepo.EXPECT().
			FindByID(gomock.Any(), gomock.Any()).
			Return(nil, userrepo.ErrNotFound)

		uc := usecase.NewGetUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.GetUserInput{ID: "22222222-2222-2222-2222-222222222222"})

		assert.Nil(t, output)
		var notFoundErr *apperrors.UserNotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})

	t.Run("異常系：空IDは UserNotFoundError (リポジトリは呼ばれない)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mocks.NewMockRepository(ctrl)
		// FindByID は呼ばれないはず

		uc := usecase.NewGetUserUseCase(mockRepo)
		output, err := uc.Execute(context.Background(), inputdto.GetUserInput{ID: ""})

		assert.Nil(t, output)
		var notFoundErr *apperrors.UserNotFoundError
		assert.ErrorAs(t, err, &notFoundErr)
	})
}
