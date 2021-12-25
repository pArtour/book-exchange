package repository

import (
	"database/sql"
	"fmt"
	"github.com/pArtour/book-exchange/pkg/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *model.User) error {
	query := fmt.Sprintf("INSERT INTO \"%s\" (email, first_name, last_name, hashed_password) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)
	err := r.db.QueryRow(
		query,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Password,
	).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetUser(email, password string) (*model.User, error) {
	user := &model.User{}
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE email=$1 AND hashed_password=$2", usersTable)
	err := r.db.QueryRow(
		query,
		email,
		password,
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
