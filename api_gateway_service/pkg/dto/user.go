package dto

type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponseDto struct {
	Uuid  string `json:"uuid"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserCredentialsRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
