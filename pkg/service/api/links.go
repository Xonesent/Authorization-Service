package service_api

import (
	"database/sql"
	"math/rand"
	"project/pkg/repository"
	"project/server"
	"time"
)

type Link_service struct {
	repos repository.Links
}

func New_Links_service(repos repository.Links) *Link_service {
	return &Link_service{repos: repos}
}

func (a *Link_service) Create_Short_URL(link *server.Link) (string, error) {
	link.Short_URL = Generate_Short_URL()
	Short_URL, err := a.repos.Create_Short_URL(link)
	if err != nil {
		return "", err
	}
	return Short_URL, nil
}

func (a *Link_service) Get_Base_URL(link *server.Link) (string, error) {
	Base_URL, err := a.repos.Get_Base_URL(link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return Base_URL, nil
}

func Generate_Short_URL() string {
	rand.Seed(time.Now().UnixNano())
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

	result := make([]byte, 10)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}
