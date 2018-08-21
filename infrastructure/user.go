package infrastructure

import (
	"context"
	"database/sql"
	"log"

	"github.com/mafuyuk/ddd-go-api-template/domain"
	"github.com/mafuyuk/ddd-go-api-template/domain/repository"
)

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r *userRepository) WithTransaction(ctx context.Context, txFunc func(*sql.Tx) error) error {
	return nil
}

func (r *userRepository) Get(domain.UserID) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) Create(*domain.User) error {
	log.Println("called infrastructure Create")
	return nil
}

func (r *userRepository) Update(*domain.User) error {
	return nil
}
