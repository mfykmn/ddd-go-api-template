package application

import (
	"context"
	"database/sql"
	"log"

	"github.com/mafuyuk/ddd-go-api-template/domain"
	"github.com/mafuyuk/ddd-go-api-template/domain/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Refer(domain.UserID) (*domain.User, error) {
	return nil, nil
}

func (s *userService) Register(ctx context.Context, user *domain.User) error {
	log.Println("called application Register")

	err := s.userRepo.WithTransaction(ctx, func(tx *sql.Tx) error {
		// insert a record into table
		if err := s.userRepo.Create(user); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *userService) Edit(ctx context.Context, user *domain.User) error {
	return nil
}
