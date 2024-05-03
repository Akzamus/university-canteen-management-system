package converter

import (
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
)

func ToUserResponseDto(user *model.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		Uuid:     user.Uuid,
		Email:    user.Email,
		Password: user.Password,
		Role:     string(user.Role),
	}
}

func ToUserModel(userDto *dto.UserRequestDto) model.User {
	return model.User{
		Email:    userDto.Email,
		Password: userDto.Password,
		Role:     model.Role(userDto.Role),
	}
}

func ToUserResponseDtoList(users *[]model.User) []dto.UserResponseDto {
	userDtoList := make([]dto.UserResponseDto, len(*users))
	for i, user := range *users {
		userDtoList[i] = ToUserResponseDto(&user)
	}
	return userDtoList
}
