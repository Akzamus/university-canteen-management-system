package dto

type AuthRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	AccessToken string `json:"accessToken"`
}
