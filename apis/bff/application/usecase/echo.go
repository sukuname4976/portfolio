package usecase

import (
	"context"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/bff/application/output-dto"
	echoentity "github.com/sukuname4976/portfolio/apis/bff/domain/entities/echo"
	pokemonentity "github.com/sukuname4976/portfolio/apis/bff/domain/entities/pokemon"
	pokemongateway "github.com/sukuname4976/portfolio/apis/bff/domain/gateway-interfaces/pokemon"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/domain/value-objects/message"
)

// EchoUseCase Echoユースケースインターフェース
type EchoUseCase interface {
	Execute(ctx context.Context, input inputdto.EchoInput) (*outputdto.EchoOutput, error)
}

// echoUseCaseImpl EchoUseCaseの実装
type echoUseCaseImpl struct {
	pokemonGateway pokemongateway.Gateway
	pokemonService *pokemonservice.Service
}

// NewEchoUseCase EchoUseCaseを生成
func NewEchoUseCase(
	pokemonGateway pokemongateway.Gateway,
	pokemonService *pokemonservice.Service,
) EchoUseCase {
	return &echoUseCaseImpl{
		pokemonGateway: pokemonGateway,
		pokemonService: pokemonService,
	}
}

// Execute ユースケースを実行
func (u *echoUseCaseImpl) Execute(ctx context.Context, input inputdto.EchoInput) (*outputdto.EchoOutput, error) {
	// 1. 入力からEchoエンティティを生成
	msg, err := message.New(input.Message)
	if err != nil {
		return nil, err
	}
	echo := echoentity.New(msg)

	// 2. ランダムなポケモンIDを生成
	pokemonID := u.pokemonService.GenerateRandomID()

	// 3. PokeAPIからポケモン情報を取得
	pokemon, err := u.pokemonGateway.FetchByID(ctx, pokemonID)
	if err != nil {
		return nil, err
	}

	// 4. OutputDTOを返却
	return u.toOutput(echo, pokemon), nil
}

// toOutput エンティティからOutputDTOに変換
func (u *echoUseCaseImpl) toOutput(echo *echoentity.Echo, pokemon *pokemonentity.Pokemon) *outputdto.EchoOutput {
	types := make([]outputdto.PokemonTypeDTO, len(pokemon.Types()))
	for i, t := range pokemon.Types() {
		types[i] = outputdto.PokemonTypeDTO{
			Slot: t.Slot(),
			Name: t.Name(),
		}
	}

	return &outputdto.EchoOutput{
		Echo: outputdto.EchoDTO{
			Message: echo.Message().Value(),
		},
		Pokemon: outputdto.PokemonDTO{
			ID:     pokemon.ID().Value(),
			Name:   pokemon.Name(),
			Types:  types,
			Sprite: pokemon.Sprite(),
		},
	}
}
