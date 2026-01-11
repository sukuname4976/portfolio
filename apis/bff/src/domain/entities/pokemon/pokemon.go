package pokemon

import (
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemonid"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemontype"
)

// Pokemon ポケモン情報を表すエンティティ
type Pokemon struct {
	id     pokemonid.PokemonID
	name   string
	types  []pokemontype.PokemonType
	sprite string
}

// New Pokemonエンティティを生成
func New(id pokemonid.PokemonID, name string, types []pokemontype.PokemonType, sprite string) *Pokemon {
	return &Pokemon{
		id:     id,
		name:   name,
		types:  types,
		sprite: sprite,
	}
}

// ID ポケモンIDを取得
func (p *Pokemon) ID() pokemonid.PokemonID {
	return p.id
}

// Name ポケモン名を取得
func (p *Pokemon) Name() string {
	return p.name
}

// Types ポケモンタイプを取得
func (p *Pokemon) Types() []pokemontype.PokemonType {
	return p.types
}

// Sprite スプライト画像URLを取得
func (p *Pokemon) Sprite() string {
	return p.sprite
}
