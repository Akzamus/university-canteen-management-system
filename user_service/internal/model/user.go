package model

type User struct {
	Uuid     string
	Email    string
	Password string
	Role     Role
}
