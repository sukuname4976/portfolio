package pokemontype

// PokemonType ポケモンタイプ値オブジェクト
type PokemonType struct {
	slot int
	name string
}

// New PokemonTypeを生成
func New(slot int, name string) PokemonType {
	return PokemonType{slot: slot, name: name}
}

// Slot スロット番号を取得
func (t PokemonType) Slot() int {
	return t.slot
}

// Name タイプ名を取得
func (t PokemonType) Name() string {
	return t.name
}
