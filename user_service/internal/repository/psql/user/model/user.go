package model

import (
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
)

type User struct {
	UUID     string     `db:"id"`
	Email    string     `db:"email"`
	Password string     `db:"password"`
	Role     model.Role `db:"role"`
}
