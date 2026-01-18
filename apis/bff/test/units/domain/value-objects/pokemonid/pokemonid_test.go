package pokemonid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/pokemonid"
)

func TestPokemonID_New(t *testing.T) {
	t.Run("境界値1でPokemonID生成成功", func(t *testing.T) {
		id, err := pokemonid.New(1)

		require.NoError(t, err)
		assert.Equal(t, 1, id.Value())
	})

	t.Run("境界値1010でPokemonID生成成功", func(t *testing.T) {
		id, err := pokemonid.New(1010)

		require.NoError(t, err)
		assert.Equal(t, 1010, id.Value())
	})

	t.Run("0でErrInvalid", func(t *testing.T) {
		_, err := pokemonid.New(0)

		assert.ErrorIs(t, err, pokemonid.ErrInvalid)
	})

	t.Run("1011でErrInvalid", func(t *testing.T) {
		_, err := pokemonid.New(1011)

		assert.ErrorIs(t, err, pokemonid.ErrInvalid)
	})

	t.Run("負の値でErrInvalid", func(t *testing.T) {
		_, err := pokemonid.New(-1)

		assert.ErrorIs(t, err, pokemonid.ErrInvalid)
	})
}
