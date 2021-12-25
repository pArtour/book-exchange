package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pArtour/book-exchange/configs"
	"github.com/pArtour/book-exchange/pkg/model"
)

type Auth interface {
	CreateUser(user *model.User) error
	GetUser(email, password string) (*model.User, error)
}

type Book interface {
}

type Repository struct {
	config *configs.Config
	Auth
	Book
}

func NewRepository(config *configs.Config, db *sql.DB) *Repository {
	return &Repository{
		config: config,
		Auth:   NewAuthRepository(db),
	}
}
