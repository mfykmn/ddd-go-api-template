package service

import (
	"github.com/mafuyuk/ddd-go-api-template/domain"
)

type UserService interface {
	Refer(domain.UserID) (*domain.User, error)
	Register(*domain.User) error
	Edit(*domain.User) error
}
