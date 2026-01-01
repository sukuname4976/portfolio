package pokemon

import (
	"math/rand"
	"time"

	"github.com/sukuname4976/portfolio/apis/bff/domain/value-objects/pokemonid"
)

// Service ポケモン取得のドメインサービス
type Service struct {
	rng *rand.Rand
}

// NewService Serviceを生成
func NewService() *Service {
	return &Service{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateRandomID ランダムなポケモンIDを生成
func (s *Service) GenerateRandomID() pokemonid.PokemonID {
	id, _ := pokemonid.New(s.rng.Intn(1010) + 1)
	return id
}
