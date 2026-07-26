package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sukuname4976/portfolio/apis/bff/src/domain/entities/user"
	usergateway "github.com/sukuname4976/portfolio/apis/bff/src/domain/gateway-interfaces/user"
)

// コンパイル時にインターフェース実装を検証
var _ usergateway.Gateway = (*UserAPIGateway)(nil)

// userServiceResponse user サービスの成功レスポンス ({ "user": {...} })
type userServiceResponse struct {
	User struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}

// userServiceErrorResponse user サービスのエラーレスポンス ({ "error": "..." })
type userServiceErrorResponse struct {
	Error string `json:"error"`
}

// createUserRequestBody user サービスへの登録リクエストボディ
type createUserRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserAPIGateway user サービス呼び出しの実装
type UserAPIGateway struct {
	baseURL    string
	httpClient *http.Client
}

// NewUserAPIGateway UserAPIGatewayを生成
func NewUserAPIGateway(baseURL string, timeout time.Duration) *UserAPIGateway {
	return &UserAPIGateway{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Create user サービスにユーザー登録を依頼する
func (g *UserAPIGateway) Create(ctx context.Context, name, email string) (*user.User, error) {
	body, err := json.Marshal(createUserRequestBody{Name: name, Email: email})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := g.baseURL + "/api/v1/users"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, g.toGatewayError(resp)
	}
	return g.decodeUser(resp)
}

// FetchByID user サービスから指定IDのユーザーを取得する
func (g *UserAPIGateway) FetchByID(ctx context.Context, id string) (*user.User, error) {
	endpoint := fmt.Sprintf("%s/api/v1/users/%s", g.baseURL, url.PathEscape(id))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, g.toGatewayError(resp)
	}
	return g.decodeUser(resp)
}

// decodeUser 成功レスポンスをドメインエンティティに変換する
func (g *UserAPIGateway) decodeUser(resp *http.Response) (*user.User, error) {
	var body userServiceResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("failed to decode user service response: %w", err)
	}
	return user.New(body.User.ID, body.User.Name, body.User.Email), nil
}

// toGatewayError 非 2xx レスポンスを GatewayError に変換する
func (g *UserAPIGateway) toGatewayError(resp *http.Response) error {
	var body userServiceErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&body)
	message := body.Error
	if message == "" {
		message = http.StatusText(resp.StatusCode)
	}
	return &usergateway.GatewayError{
		StatusCode: resp.StatusCode,
		Message:    message,
	}
}
