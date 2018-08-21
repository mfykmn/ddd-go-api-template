package service

import (
	"context"

	"github.com/mafuyuk/ddd-go-api-template/domain"
)

type UserService interface {
	Refer(domain.UserID) (*domain.User, error)
	Register(context.Context, *domain.User) error
	Edit(context.Context, *domain.User) error
}
