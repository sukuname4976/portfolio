package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/pokemon"
	pokemongateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemonid"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemontype"
)

// コンパイル時にインターフェース実装を検証
var _ pokemongateway.Gateway = (*PokeAPIGateway)(nil)

// pokeAPIResponse PokeAPIのレスポンス構造体
type pokeAPIResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

// PokeAPIGateway PokeAPI呼び出しの実装
type PokeAPIGateway struct {
	baseURL    string
	httpClient *http.Client
}

// NewPokeAPIGateway PokeAPIGatewayを生成
func NewPokeAPIGateway(baseURL string, timeout time.Duration) *PokeAPIGateway {
	return &PokeAPIGateway{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// FetchByID PokeAPIからポケモン情報を取得
func (g *PokeAPIGateway) FetchByID(ctx context.Context, id pokemonid.PokemonID) (*pokemon.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%d", g.baseURL, id.Value())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pokemon not found: status %d", resp.StatusCode)
	}

	var apiResp pokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return g.toDomain(apiResp)
}

// toDomain APIレスポンスをドメインエンティティに変換
func (g *PokeAPIGateway) toDomain(resp pokeAPIResponse) (*pokemon.Pokemon, error) {
	id, err := pokemonid.New(resp.ID)
	if err != nil {
		return nil, err
	}

	types := make([]pokemontype.PokemonType, len(resp.Types))
	for i, t := range resp.Types {
		types[i] = pokemontype.New(t.Slot, t.Type.Name)
	}

	return pokemon.New(id, resp.Name, types, resp.Sprites.FrontDefault), nil
}
