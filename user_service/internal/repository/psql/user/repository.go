package user

import (
	"context"
	"fmt"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql/user/converter"
	"github.com/jmoiron/sqlx"

	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/model"
	def "github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository"
	repoModel "github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/repository/psql/user/model"
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

func (r *repository) FindById(_ context.Context, uuid string) (*model.User, error) {
	var user repoModel.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", psql.UserTable)

	if err := r.db.Get(&user, query, uuid); err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) FindAll(_ context.Context) (*[]model.User, error) {
	var users []repoModel.User
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %1", psql.UserTable)

	if err := r.db.Select(&users, query, rowLimit); err != nil {
		return nil, err
	}

	return converter.ToUsersFromRepo(&users), nil
}

func (r *repository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	userRepo := converter.ToUserFromService(user)

	query := fmt.Sprintf("INSERT INTO %s (email, password, role) VALUES ($1, $2, $3) RETURNING *", psql.UserTable)
	params := []interface{}{userRepo.Email, userRepo.Password, userRepo.Role}

	if userRepo.UUID != "" {
		query = fmt.Sprintf("UPDATE %s SET email = $1, password = $2, role = $3 WHERE id = $4 RETURNING *", psql.UserTable)
		params = append(params, userRepo.UUID)
	}

	if err := r.db.GetContext(ctx, userRepo, query, params...); err != nil {
		return nil, err
	}
	return converter.ToUserFromRepo(userRepo), nil
}

func (r *repository) DeleteById(ctx context.Context, uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", psql.UserTable)
	_, err := r.db.ExecContext(ctx, query, uuid)
	return err
}
