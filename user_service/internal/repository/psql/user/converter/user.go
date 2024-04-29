package converter

import (
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
	repoModel "github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql/user/model"
)

func ToUser(user *repoModel.User) model.User {
	return model.User{
		UUID:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToUsers(users *[]repoModel.User) []model.User {
	servUsers := make([]model.User, len(*users))
	for i, user := range *users {
		servUsers[i] = ToUser(&user)
	}
	return servUsers
}

func ToRepoUser(user *model.User) repoModel.User {
	return repoModel.User{
		UUID:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
