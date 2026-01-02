package outputdto

// EchoOutput Echoユースケースの出力
type EchoOutput struct {
	Echo    EchoDTO
	Pokemon PokemonDTO
}

// EchoDTO エコー情報のDTO
type EchoDTO struct {
	Message string
}

// PokemonDTO ポケモン情報のDTO
type PokemonDTO struct {
	ID     int
	Name   string
	Types  []PokemonTypeDTO
	Sprite string
}

// PokemonTypeDTO ポケモンタイプのDTO
type PokemonTypeDTO struct {
	Slot int
	Name string
}
