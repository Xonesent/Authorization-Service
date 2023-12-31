package repository_api

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Auth_repository struct {
	db *sqlx.DB
}

func New_Auth_repository(db *sqlx.DB) *Auth_repository {
	return &Auth_repository{db: db}
}

func (r *Auth_repository) Create_User(login string, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) values ($1, $2) RETURNING id", Users_Table)

	row := r.db.QueryRow(query, login, password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Auth_repository) Get_User(login string, password string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", Users_Table)
	err := r.db.Get(&id, query, login, password)
	
	return id, err
}