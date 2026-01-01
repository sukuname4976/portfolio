package responsemodel

// EchoResponse /echo レスポンスボディ
type EchoResponse struct {
	Echo    EchoData    `json:"echo"`
	Pokemon PokemonData `json:"pokemon"`
}

// EchoData エコー情報
type EchoData struct {
	Message string `json:"message"`
}

// PokemonData ポケモン情報
type PokemonData struct {
	ID     int        `json:"id"`
	Name   string     `json:"name"`
	Types  []TypeData `json:"types"`
	Sprite string     `json:"sprite"`
}

// TypeData タイプ情報
type TypeData struct {
	Slot int    `json:"slot"`
	Name string `json:"name"`
}

// ErrorResponse エラーレスポンス
type ErrorResponse struct {
	Error string `json:"error"`
}
