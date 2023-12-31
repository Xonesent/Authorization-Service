package service

import (
	"project/pkg/repository"
	service_api "project/pkg/service/api"
	"project/server"
)

type Authorization interface{
	Create_User(login string, password string) (int, error)
	Generate_Token(login string, password string) (string, error)
	Parse_Token(token string) (int, error)
}

type Links interface {
	Create_Short_URL(*server.Link) (string, error)
	Get_Base_URL(*server.Link) (string, error)
}

type Service struct {
	Authorization
	Links
}

func New_Service(repos *repository.Repository) *Service {
	return &Service{
		Authorization: service_api.New_Auth_service(repos.Authorization),
		Links: service_api.New_Links_service(repos.Links),
	}
}