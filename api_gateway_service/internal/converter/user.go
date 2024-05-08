package converter

import (
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/model"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
)

func ToUserRequestDto(auth *dto.AuthRequestDto) dto.UserRequestDto {
	return dto.UserRequestDto{
		Email:    auth.Email,
		Password: auth.Password,
		Role:     string(model.UserRole),
	}
}

func ToUserCredentialsRequestDto(auth *dto.AuthRequestDto) dto.UserCredentialsRequestDto {
	return dto.UserCredentialsRequestDto{
		Email:    auth.Email,
		Password: auth.Password,
	}
}
