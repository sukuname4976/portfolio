package controller

import (
	"context"
	"strings"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/application/input-dto"
	"github.com/sukuname4976/portfolio/apis/bff/application/usecase"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/config"
	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/gateway"
	ogen "github.com/sukuname4976/portfolio/apis/bff/presentation/auto-generated-by-ogen"
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
		return &ogen.EchoBadRequest{
			Error: "message is required",
		}, nil
	}

	// ユースケースの実行
	input := inputdto.EchoInput{
		Message: req.Message,
	}
	output, err := h.echoUseCase.Execute(ctx, input)
	if err != nil {
		return &ogen.EchoInternalServerError{
			Error: err.Error(),
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
