package service

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, uuid string) (dto.UserResponseDto, error)
	GetAllUsers(ctx context.Context) ([]dto.UserResponseDto, error)
	CreateUser(ctx context.Context, userDto dto.UserRequestDto) (dto.UserResponseDto, error)
	UpdateUser(ctx context.Context, userDto dto.UserRequestDto, uuid string) (dto.UserResponseDto, error)
	DeleteUserByID(ctx context.Context, uuid string) error
	VerifyUserCredentials(ctx context.Context, userCredentialsDto dto.UserCredentialsRequestDto) (dto.UserResponseDto, error)
}

type AuthService interface {
	Register(ctx context.Context, authDto dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Authenticate(ctx context.Context, authDto dto.AuthRequestDto) (dto.AuthResponseDto, error)
}
