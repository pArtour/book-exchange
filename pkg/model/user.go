package model

type User struct {
	ID        int    `json:"id" db:"id"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
