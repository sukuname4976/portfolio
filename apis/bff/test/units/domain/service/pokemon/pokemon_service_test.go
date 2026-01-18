package pokemon_service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pokemonservice "github.com/sukuname4976/portfolio/apis/bff/src/domain/service/pokemon"
)

func TestPokemonService_GenerateRandomID(t *testing.T) {
	t.Run("生成されるIDが1-1010の範囲内", func(t *testing.T) {
		svc := pokemonservice.NewService()

		// 複数回実行して範囲内か確認
		for i := 0; i < 100; i++ {
			id := svc.GenerateRandomID()
			value := id.Value()

			assert.GreaterOrEqual(t, value, 1)
			assert.LessOrEqual(t, value, 1010)
		}
	})
}
