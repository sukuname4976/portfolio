package user

import (
	"context"
	"fmt"

	"github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/user"
)

// Gateway user サービス呼び出しインターフェース
type Gateway interface {
	// Create user サービスにユーザー登録を依頼し、登録済みユーザーを返す
	Create(ctx context.Context, name, email string) (*user.User, error)
	// FetchByID user サービスから指定IDのユーザーを取得する
	FetchByID(ctx context.Context, id string) (*user.User, error)
}

// GatewayError user サービスが 2xx 以外を返した際に、上流のステータスコードと
// メッセージを保持して返すエラー。controller でステータスをマッピングするために使う。
type GatewayError struct {
	StatusCode int
	Message    string
}

// Error errorインターフェースを実装
func (e *GatewayError) Error() string {
	return fmt.Sprintf("user service responded with status %d: %s", e.StatusCode, e.Message)
}
