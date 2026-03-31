package userid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
)

func TestUserID_New(t *testing.T) {
	t.Run("正常な文字列でUserID生成成功", func(t *testing.T) {
		id, err := userid.New("user-123")

		require.NoError(t, err)
		assert.Equal(t, "user-123", id.Value())
	})

	t.Run("空文字でErrEmpty", func(t *testing.T) {
		id, err := userid.New("")

		assert.ErrorIs(t, err, userid.ErrEmpty)
		assert.Equal(t, "", id.Value())
	})
}
