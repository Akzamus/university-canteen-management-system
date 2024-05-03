package service

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, uuid string) (dto.UserResponseDto, error)
	GetAllUsers(ctx context.Context) ([]dto.UserResponseDto, error)
	CreateUser(ctx context.Context, user dto.UserRequestDto) (dto.UserResponseDto, error)
	UpdateUser(ctx context.Context, user dto.UserRequestDto, uuid string) (dto.UserResponseDto, error)
	DeleteUserByID(ctx context.Context, uuid string) error
}
