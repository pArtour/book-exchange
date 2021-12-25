package service

import (
	"github.com/pArtour/book-exchange/pkg/model"
	"github.com/pArtour/book-exchange/pkg/repository"
)

type Auth interface {
	CreateUser(user *model.User) error
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Book interface {
}

type Service struct {
	Auth
	Book
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository),
	}
}
