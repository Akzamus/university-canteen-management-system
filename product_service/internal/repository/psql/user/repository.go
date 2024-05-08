package user

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/model"
	def "github.com/Akzamus/university-canteen-management-system/user_service/internal/repository"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository/psql/user/converter"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository/psql/user/entity"
	"github.com/jmoiron/sqlx"
)

const (
	rowLimit = 1000
)

var _ def.ProductRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindById(_ context.Context, uuid string) (model.Product, error) {
	product := entity.Product{}
	query := "SELECT * FROM _product WHERE id = $1"

	err := r.db.Get(&product, query, uuid)
	return converter.ToProductModel(&product), err
}

func (r *repository) FindAll(_ context.Context) ([]model.Product, error) {
	var product []entity.User
	query := "SELECT * FROM _product LIMIT $1"

	err := r.db.Select(&product, query, rowLimit)
	return converter.ToProductModelList(&product), err
}

func (r *repository) Save(ctx context.Context, product model.Product) (model.Product, error) {
	ProductRepo := converter.ToProductEntity(&product)

	query := "INSERT INTO _product (name, description, price, image_url, category_id, total_quantity) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	params := []interface{}{ProductRepo.Name, ProductRepo.Description, ProductRepo.Price, ProductRepo.Image_url, ProductRepo.category_id, ProductRepo.total_quantity}

	if ProductRepo.Uuid != "" {
		query = "UPDATE _product SET name = $1, description = $2, price = $3, image_url = $4, category_id = $5, total_quantity = $6 WHERE id = $7 RETURNING *"
		params = append(params, product.Uuid)
	}

	err := r.db.GetContext(ctx, &ProductRepo, query, params...)
	return converter.ToUserModel(&ProductRepo), err
}

func (r *repository) DeleteById(ctx context.Context, uuid string) error {
	query := "DELETE FROM _product WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	return err
}
