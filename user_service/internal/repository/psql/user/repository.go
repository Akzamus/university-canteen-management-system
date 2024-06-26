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

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindById(_ context.Context, uuid string) (model.User, error) {
	user := entity.User{}
	query := "SELECT * FROM _user WHERE id = $1"

	err := r.db.Get(&user, query, uuid)
	return converter.ToUserModel(&user), err
}

func (r *repository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	user := entity.User{}
	query := "SELECT * FROM _user WHERE email = $1"

	err := r.db.GetContext(ctx, &user, query, email)
	return converter.ToUserModel(&user), err
}

func (r *repository) FindAll(_ context.Context) ([]model.User, error) {
	var users []entity.User
	query := "SELECT * FROM _user LIMIT $1"

	err := r.db.Select(&users, query, rowLimit)
	return converter.ToUserModelList(&users), err
}

func (r *repository) Save(ctx context.Context, user model.User) (model.User, error) {
	userRepo := converter.ToUserEntity(&user)

	query := "INSERT INTO _user (email, password, role) VALUES ($1, $2, $3) RETURNING *"
	params := []interface{}{userRepo.Email, userRepo.Password, userRepo.Role}

	if userRepo.Uuid != "" {
		query = "UPDATE _user SET email = $1, password = $2, role = $3 WHERE id = $4 RETURNING *"
		params = append(params, userRepo.Uuid)
	}

	err := r.db.GetContext(ctx, &userRepo, query, params...)
	return converter.ToUserModel(&userRepo), err
}

func (r *repository) DeleteById(ctx context.Context, uuid string) error {
	query := "DELETE FROM _user WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	return err
}
