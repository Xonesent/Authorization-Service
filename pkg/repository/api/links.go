package repository_api

import (
	"fmt"
	"project/server"

	"github.com/jmoiron/sqlx"
)

type Links_repository struct {
	db *sqlx.DB
}

func New_Links_repository(db *sqlx.DB) *Links_repository {
	return &Links_repository{db: db}
}

func (a *Links_repository) Create_Short_URL(link *server.Link) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (base_url, short_url) values ($1, $2)", Links_Table)
	a.db.QueryRow(query, link.Base_URL, link.Short_URL)

	return link.Short_URL, nil
}

func (a *Links_repository) Get_Base_URL(link *server.Link) (string, error) {
	var base_url string
	query := fmt.Sprintf("SELECT base_url FROM %s WHERE short_url=$1", Links_Table)
	err := a.db.Get(&base_url, query, link.Short_URL)

	return base_url, err
}
