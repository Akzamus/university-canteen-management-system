package model

type User struct {
	UUID     string
	Email    string
	Password string
	Role     Role
}
