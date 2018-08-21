package repository

import (
	"context"
	"database/sql"

	"github.com/mafuyuk/ddd-go-api-template/domain"
)

type UserRepository interface {
	WithTransaction(ctx context.Context, txFunc func(*sql.Tx) error) error
	Get(domain.UserID) (*domain.User, error)
	Create(*domain.User) error
	Update(*domain.User) error
}
