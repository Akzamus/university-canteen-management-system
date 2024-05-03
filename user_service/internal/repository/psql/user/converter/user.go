package converter

import (
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql/user/entity"
)

func ToModel(user *entity.User) model.User {
	return model.User{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToModels(users *[]entity.User) []model.User {
	servUsers := make([]model.User, len(*users))
	for i, user := range *users {
		servUsers[i] = ToModel(&user)
	}
	return servUsers
}

func ToEntity(user *model.User) entity.User {
	return entity.User{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
