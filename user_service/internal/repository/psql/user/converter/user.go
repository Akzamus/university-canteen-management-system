package converter

import (
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
	repoModel "github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql/user/model"
)

func ToUserFromRepo(user *repoModel.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		UUID:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToUsersFromRepo(users *[]repoModel.User) *[]model.User {
	if users == nil {
		return nil
	}

	servUsers := make([]model.User, len(*users))
	for i, user := range *users {
		servUsers[i] = *ToUserFromRepo(&user)
	}
	return &servUsers
}

func ToUserFromService(user *model.User) *repoModel.User {
	if user == nil {
		return nil
	}

	return &repoModel.User{
		UUID:     user.UUID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}
