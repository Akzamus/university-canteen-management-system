package psql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgreSqlClient(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SslMode,
		),
	)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
