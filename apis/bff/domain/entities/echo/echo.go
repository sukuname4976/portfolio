package echo

import "github.com/sukuname4976/portfolio/apis/bff/domain/value-objects/message"

// Echo エコーメッセージを表すエンティティ
type Echo struct {
	message message.Message
}

// New Echoエンティティを生成
func New(msg message.Message) *Echo {
	return &Echo{message: msg}
}

// Message メッセージを取得
func (e *Echo) Message() message.Message {
	return e.message
}
