package service

import (
	"context"

	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
)

type UserService interface {
	GetUserByID(ctx context.Context, uuid string) (dto.ProductResponseDto, error)
	GetAllUsers(ctx context.Context) ([]dto.ProductResponseDto, error)
	CreateUser(ctx context.Context, user dto.ProductRequestDto) (dto.ProductResponseDto, error)
	UpdateUser(ctx context.Context, user dto.ProductRequestDto, uuid string) (dto.ProductResponseDto, error)
	DeleteUserByID(ctx context.Context, uuid string) error
}
