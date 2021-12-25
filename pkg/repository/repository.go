package repository

type Auth interface {
}

type Record interface {
}

type Repository struct {
	Auth
	Record
}

func NewRepository() *Repository {
	return &Repository{}
}
