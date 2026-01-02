package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sukuname4976/portfolio/apis/bff/src/domain/value-objects/message"
)

func TestMessage_New(t *testing.T) {
	t.Run("正常な文字列でMessage生成成功", func(t *testing.T) {
		msg, err := message.New("hello")

		require.NoError(t, err)
		assert.Equal(t, "hello", msg.Value())
	})

	t.Run("空文字でErrEmpty", func(t *testing.T) {
		msg, err := message.New("")

		assert.ErrorIs(t, err, message.ErrEmpty)
		assert.Equal(t, "", msg.Value())
	})
}
