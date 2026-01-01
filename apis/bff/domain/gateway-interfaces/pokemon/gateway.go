package pokemon

import (
	"context"

	"github.com/sukuname4976/portfolio/apis/bff/domain/entities/pokemon"
	"github.com/sukuname4976/portfolio/apis/bff/domain/value-objects/pokemonid"
)

// Gateway 外部ポケモンAPI呼び出しインターフェース
type Gateway interface {
	// FetchByID 指定IDのポケモン情報を取得
	FetchByID(ctx context.Context, id pokemonid.PokemonID) (*pokemon.Pokemon, error)
}
