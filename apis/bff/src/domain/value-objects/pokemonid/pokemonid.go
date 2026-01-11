package pokemonid

import "errors"

const (
	minID = 1
	maxID = 1010
)

// ErrInvalid ポケモンIDが範囲外の場合のエラー
var ErrInvalid = errors.New("pokemon id must be between 1 and 1010")

// PokemonID ポケモンID値オブジェクト
type PokemonID struct {
	value int
}

// New PokemonIDを生成（1-1010の範囲チェック）
func New(value int) (PokemonID, error) {
	if value < minID || value > maxID {
		return PokemonID{}, ErrInvalid
	}
	return PokemonID{value: value}, nil
}

// Value 値を取得
func (p PokemonID) Value() int {
	return p.value
}
