package entity

import (
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/model"
)

type User struct {
	Uuid     string     `db:"id"`
	Email    string     `db:"email"`
	Password string     `db:"password"`
	Role     model.Role `db:"role"`
}
