package service

import "github.com/pArtour/book-exchange/pkg/repository"

type Auth interface {
}

type Record interface {
}

type Service struct {
	Auth
	Record
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
