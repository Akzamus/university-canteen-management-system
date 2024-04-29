package repository

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
)

type UserRepository interface {
	FindById(ctx context.Context, uuid string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Save(ctx context.Context, user model.User) (model.User, error)
	DeleteById(ctx context.Context, uuid string) error
}
