package converter

import (
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/model"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository/psql/user/entity"
)

func ToUserModel(user *entity.User) model.User {
	return model.User{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToUserModelList(users *[]entity.User) []model.User {
	servUsers := make([]model.User, len(*users))
	for i, user := range *users {
		servUsers[i] = ToUserModel(&user)
	}
	return servUsers
}

func ToUserEntity(user *model.User) entity.User {
	return entity.User{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
