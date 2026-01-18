package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/bff/src/application/usecase"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/pokemon"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/src/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/message"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemonid"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemontype"
	"github.com/sukuname4976/portfolio/apis/bff/test/mocks"
	"go.uber.org/mock/gomock"
)

func TestEchoUseCase_Execute(t *testing.T) {
	t.Run("正常系：メッセージとポケモン情報を返す", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// モックの準備
		mockGateway := mocks.NewMockGateway(ctrl)
		svc := pokemonservice.NewService()

		// テスト用ポケモンデータ
		pokemonID, _ := pokemonid.New(25)
		pikachu := pokemon.New(
			pokemonID,
			"pikachu",
			[]pokemontype.PokemonType{pokemontype.New(1, "electric")},
			"https://example.com/pikachu.png",
		)

		// モックの期待値設定（任意のIDでpikachuを返す）
		mockGateway.EXPECT().
			FetchByID(gomock.Any(), gomock.Any()).
			Return(pikachu, nil)

		// テスト実行
		uc := usecase.NewEchoUseCase(mockGateway, svc)
		input := inputdto.EchoInput{Message: "hello"}
		output, err := uc.Execute(context.Background(), input)

		// 検証
		require.NoError(t, err)
		assert.Equal(t, "hello", output.Echo.Message)
		assert.Equal(t, "pikachu", output.Pokemon.Name)
		assert.Equal(t, 25, output.Pokemon.ID)
	})

	t.Run("異常系：空メッセージでエラー", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGateway := mocks.NewMockGateway(ctrl)
		svc := pokemonservice.NewService()

		// Gatewayは呼ばれないはず

		uc := usecase.NewEchoUseCase(mockGateway, svc)
		input := inputdto.EchoInput{Message: ""}
		output, err := uc.Execute(context.Background(), input)

		assert.ErrorIs(t, err, message.ErrEmpty)
		assert.Nil(t, output)
	})

	t.Run("異常系：Gatewayエラー", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockGateway := mocks.NewMockGateway(ctrl)
		svc := pokemonservice.NewService()

		gatewayErr := errors.New("gateway error")
		mockGateway.EXPECT().
			FetchByID(gomock.Any(), gomock.Any()).
			Return(nil, gatewayErr)

		uc := usecase.NewEchoUseCase(mockGateway, svc)
		input := inputdto.EchoInput{Message: "hello"}
		output, err := uc.Execute(context.Background(), input)

		assert.ErrorIs(t, err, gatewayErr)
		assert.Nil(t, output)
	})
}
