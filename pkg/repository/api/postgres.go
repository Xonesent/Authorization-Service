package repository_api

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	Users_Table = "users"
	Links_Table = "links"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	SSLMode  string
}

func New_Postgres_DB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
