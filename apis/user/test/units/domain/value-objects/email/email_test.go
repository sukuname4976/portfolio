package email_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
)

func TestEmail_New(t *testing.T) {
	t.Run("正常なメールアドレスでEmail生成成功", func(t *testing.T) {
		e, err := email.New("test@example.com")

		require.NoError(t, err)
		assert.Equal(t, "test@example.com", e.Value())
	})

	t.Run("空文字でErrEmpty", func(t *testing.T) {
		e, err := email.New("")

		assert.ErrorIs(t, err, email.ErrEmpty)
		assert.Equal(t, "", e.Value())
	})

	t.Run("@がない文字列でErrInvalidFormat", func(t *testing.T) {
		e, err := email.New("invalid-email")

		assert.ErrorIs(t, err, email.ErrInvalidFormat)
		assert.Equal(t, "", e.Value())
	})
}
