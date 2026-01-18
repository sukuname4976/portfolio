package controller

import (
	"context"
	"strings"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/src/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/bff/src/application/usecase"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/src/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/src/infrastructure/config"
	apperrors "github.com/sukuname4976/portfolio/apis/bff/src/infrastructure/errors"
	"github.com/sukuname4976/portfolio/apis/bff/src/infrastructure/gateway"
	ogen "github.com/sukuname4976/portfolio/apis/bff/src/presentation/auto-generated-by-ogen"
	"github.com/sukuname4976/portfolio/apis/bff/src/presentation/middleware"
)

// Handler ogenのHandlerインターフェースを実装
type Handler struct {
	echoUseCase usecase.EchoUseCase
}

// NewHandler Handlerを生成
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		echoUseCase: newEchoUseCase(cfg),
	}
}

// newEchoUseCase Echoエンドポイント用のDI
func newEchoUseCase(cfg *config.Config) usecase.EchoUseCase {
	pokeAPIGateway := gateway.NewPokeAPIGateway(cfg.PokeAPIBaseURL, cfg.HTTPTimeout)
	pokemonSvc := pokemonservice.NewService()
	return usecase.NewEchoUseCase(pokeAPIGateway, pokemonSvc)
}

// Echo POST /echo を処理
func (h *Handler) Echo(ctx context.Context, req *ogen.EchoRequest) (ogen.EchoRes, error) {
	// バリデーション
	if strings.TrimSpace(req.Message) == "" {
		appErr := apperrors.NewEchoValidationError("message is required", "empty message received")
		middleware.LogError(appErr)
		return &ogen.EchoBadRequest{
			Error: appErr.ExternalErrorMessage,
		}, nil
	}

	// ユースケースの実行
	input := inputdto.EchoInput{
		Message: req.Message,
	}
	output, err := h.echoUseCase.Execute(ctx, input)
	if err != nil {
		appErr := apperrors.NewPokemonGatewayError(err.Error())
		middleware.LogError(appErr)
		return &ogen.EchoInternalServerError{
			Error: appErr.ExternalErrorMessage,
		}, nil
	}

	// レスポンスの構築
	types := make([]ogen.TypeData, len(output.Pokemon.Types))
	for i, t := range output.Pokemon.Types {
		types[i] = ogen.TypeData{
			Slot: t.Slot,
			Name: t.Name,
		}
	}

	return &ogen.EchoResponse{
		Echo: ogen.EchoData{
			Message: output.Echo.Message,
		},
		Pokemon: ogen.PokemonData{
			ID:     output.Pokemon.ID,
			Name:   output.Pokemon.Name,
			Types:  types,
			Sprite: output.Pokemon.Sprite,
		},
	}, nil
}
