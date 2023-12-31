package service_api

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"project/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt = "kjkgjrnhlmuaio412879"
	key  = "qwefgh21835!#dfks,I"
)

type Token_Claims struct {
	jwt.StandardClaims
	User_Id int `json:"user_id"`
}

type Auth_service struct {
	repos repository.Authorization
}

func New_Auth_service(repos repository.Authorization) *Auth_service {
	return &Auth_service{repos: repos}
}

func (s *Auth_service) Create_User(login string, password string) (int, error) {
	password = Generate_Hash_Password(password)
	return s.repos.Create_User(login, password)
}

func (s *Auth_service) Generate_Token(login string, password string) (string, error) {
	id, err := s.repos.Get_User(login, Generate_Hash_Password(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token_Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(key))
}

func (s *Auth_service) Parse_Token(token string) (int, error) {
	validated_token, err := jwt.ParseWithClaims(token, &Token_Claims{}, func(validated_token *jwt.Token) (interface{}, error) {
		if _, ok := validated_token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := validated_token.Claims.(*Token_Claims)
	if !ok {
		return 0, errors.New("token claims are not of type *Token_Claims")
	}

	return claims.User_Id, nil
}

func Generate_Hash_Password(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
