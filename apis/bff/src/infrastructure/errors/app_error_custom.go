package errors

import "net/http"

// EchoValidationError echoエンドポイントのバリデーションエラー
type EchoValidationError struct {
	BadRequestError
}

// NewEchoValidationError echoバリデーションエラーを生成
func NewEchoValidationError(externalMsg, internalMsg string) *EchoValidationError {
	return &EchoValidationError{
		BadRequestError: BadRequestError{
			BaseError: BaseError{
				ErrorID:              ErrIDEchoValidation,
				Status:               http.StatusBadRequest,
				ExternalErrorMessage: externalMsg,
				InternalErrorMessage: internalMsg,
			},
		},
	}
}

// PokemonGatewayError PokeAPI通信エラー
type PokemonGatewayError struct {
	BadGatewayError
}

// NewPokemonGatewayError PokeAPI通信エラーを生成
func NewPokemonGatewayError(internalMsg string) *PokemonGatewayError {
	return &PokemonGatewayError{
		BadGatewayError: BadGatewayError{
			BaseError: BaseError{
				ErrorID:              ErrIDPokemonGateway,
				Status:               http.StatusBadGateway,
				ExternalErrorMessage: "pokemon service unavailable",
				InternalErrorMessage: internalMsg,
			},
		},
	}
}

// UserGatewayError user サービス通信エラー
type UserGatewayError struct {
	BadGatewayError
}

// NewUserGatewayError user サービス通信エラーを生成
func NewUserGatewayError(internalMsg string) *UserGatewayError {
	return &UserGatewayError{
		BadGatewayError: BadGatewayError{
			BaseError: BaseError{
				ErrorID:              ErrIDUserGateway,
				Status:               http.StatusBadGateway,
				ExternalErrorMessage: "user service unavailable",
				InternalErrorMessage: internalMsg,
			},
		},
	}
}
