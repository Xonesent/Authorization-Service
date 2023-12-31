package repository

import (
	repository_api "project/pkg/repository/api"
	"project/server"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create_User(login string, password string) (int, error)
	Get_User(login string, password string) (int, error)
}

type Links interface {
	Create_Short_URL(*server.Link) (string, error)
	Get_Base_URL(*server.Link) (string, error)
}

type Repository struct {
	Authorization
	Links
}

func New_Repository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: repository_api.New_Auth_repository(db),
		Links: repository_api.New_Links_repository(db),
	}
}
