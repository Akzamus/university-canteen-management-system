package repository

import (
	"context"

	"github.com/Akzamus/university-canteen-management-system/product_service/internal/model"
)

type ProductRepository interface {
	FindById(ctx context.Context, uuid string) (model.Product, error)
	FindAll(ctx context.Context) ([]model.Product, error)
	Save(ctx context.Context, product model.Product) (model.Product, error)
	DeleteById(ctx context.Context, uuid string) error
}
