package inputdto

// CreateUserInput ユーザー登録ユースケースの入力DTO
type CreateUserInput struct {
	Name  string
	Email string
}

// GetUserInput ユーザー取得ユースケースの入力DTO
type GetUserInput struct {
	ID string
}
