package outputdto

// GetUserOutput GetUserユースケースの出力DTO
type GetUserOutput struct {
	User UserDTO
}

// UserDTO ユーザー情報DTO
type UserDTO struct {
	ID    string
	Name  string
	Email string
}
