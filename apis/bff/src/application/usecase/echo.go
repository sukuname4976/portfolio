package usecase

import (
	"context"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/output-dto"
	pokemonentity "github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/pokemon"
	pokemongateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/pokemon"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/src/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/message"
)

// コンパイル時にインターフェース実装を検証
var _ EchoUseCase = (*echoUseCase)(nil)

// EchoUseCase Echoユースケースインターフェース
type EchoUseCase interface {
	Execute(ctx context.Context, input inputdto.EchoInput) (*outputdto.EchoOutput, error)
}

// echoUseCase EchoUseCaseの実装
type echoUseCase struct {
	pokemonGateway pokemongateway.Gateway
	pokemonService *pokemonservice.Service
}

// NewEchoUseCase EchoUseCaseを生成
func NewEchoUseCase(
	pokemonGateway pokemongateway.Gateway,
	pokemonService *pokemonservice.Service,
) EchoUseCase {
	return &echoUseCase{
		pokemonGateway: pokemonGateway,
		pokemonService: pokemonService,
	}
}

// Execute ユースケースを実行
func (u *echoUseCase) Execute(ctx context.Context, input inputdto.EchoInput) (*outputdto.EchoOutput, error) {
	// 1. 入力メッセージをバリデーション
	msg, err := message.New(input.Message)
	if err != nil {
		return nil, err
	}

	// 2. ランダムなポケモンIDを生成
	pokemonID := u.pokemonService.GenerateRandomID()

	// 3. PokeAPIからポケモン情報を取得
	pokemon, err := u.pokemonGateway.FetchByID(ctx, pokemonID)
	if err != nil {
		return nil, err
	}

	// 4. OutputDTOを返却
	return u.toOutput(msg, pokemon), nil
}

// toOutput ドメインオブジェクトからOutputDTOに変換
func (u *echoUseCase) toOutput(msg message.Message, pokemon *pokemonentity.Pokemon) *outputdto.EchoOutput {
	types := make([]outputdto.PokemonTypeDTO, len(pokemon.Types()))
	for i, t := range pokemon.Types() {
		types[i] = outputdto.PokemonTypeDTO{
			Slot: t.Slot(),
			Name: t.Name(),
		}
	}

	return &outputdto.EchoOutput{
		Echo: outputdto.EchoDTO{
			Message: msg.Value(),
		},
		Pokemon: outputdto.PokemonDTO{
			ID:     pokemon.ID().Value(),
			Name:   pokemon.Name(),
			Types:  types,
			Sprite: pokemon.Sprite(),
		},
	}
}
