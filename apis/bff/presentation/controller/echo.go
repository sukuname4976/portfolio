package controller

import (
	"encoding/json"
	"net/http"

	inputdto "github.com/sukuname4976/portfolio/apis/bff/application/input-dto"
	outputdto "github.com/sukuname4976/portfolio/apis/bff/application/output-dto"
	"github.com/sukuname4976/portfolio/apis/bff/application/usecase"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/domain/service/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/config"
	"github.com/sukuname4976/portfolio/apis/bff/infrastructure/gateway"
	requestmodel "github.com/sukuname4976/portfolio/apis/bff/presentation/request-model"
	responsemodel "github.com/sukuname4976/portfolio/apis/bff/presentation/response-model"
)

// EchoController Echoエンドポイントのコントローラー
type EchoController struct {
	echoUseCase usecase.EchoUseCase
}

// NewEchoController EchoControllerを生成
func NewEchoController(echoUseCase usecase.EchoUseCase) *EchoController {
	return &EchoController{
		echoUseCase: echoUseCase,
	}
}

// NewEchoHandler Echoエンドポイントのハンドラーを生成（DI込み）
func NewEchoHandler(cfg *config.Config) http.HandlerFunc {
	// DI: このエンドポイントに必要な依存関係をここで解決
	pokeAPIGateway := gateway.NewPokeAPIGateway(cfg.PokeAPIBaseURL, cfg.HTTPTimeout)
	pokemonSvc := pokemonservice.NewService()
	echoUseCase := usecase.NewEchoUseCase(pokeAPIGateway, pokemonSvc)
	ctrl := NewEchoController(echoUseCase)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		ctrl.Handle(w, r)
	}
}

// Handle POST /echo を処理
func (c *EchoController) Handle(w http.ResponseWriter, r *http.Request) {
	// リクエストのデコード
	var req requestmodel.EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.respondError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// バリデーション
	if err := req.Validate(); err != nil {
		c.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ユースケースの実行
	input := inputdto.EchoInput{
		Message: req.Message,
	}
	output, err := c.echoUseCase.Execute(r.Context(), input)
	if err != nil {
		c.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// レスポンスの構築
	response := c.toResponse(output)
	c.respondJSON(w, response, http.StatusOK)
}

// toResponse OutputDTOからレスポンスモデルに変換
func (c *EchoController) toResponse(output *outputdto.EchoOutput) *responsemodel.EchoResponse {
	types := make([]responsemodel.TypeData, len(output.Pokemon.Types))
	for i, t := range output.Pokemon.Types {
		types[i] = responsemodel.TypeData{
			Slot: t.Slot,
			Name: t.Name,
		}
	}

	return &responsemodel.EchoResponse{
		Echo: responsemodel.EchoData{
			Message: output.Echo.Message,
		},
		Pokemon: responsemodel.PokemonData{
			ID:     output.Pokemon.ID,
			Name:   output.Pokemon.Name,
			Types:  types,
			Sprite: output.Pokemon.Sprite,
		},
	}
}

// respondJSON JSONレスポンスを返す
func (c *EchoController) respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError エラーレスポンスを返す
func (c *EchoController) respondError(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responsemodel.ErrorResponse{Error: message})
}
