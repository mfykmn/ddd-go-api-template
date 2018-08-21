package application

import (
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

func (s *userService) Register(user *domain.User) error {
	log.Println("called application Register")
	s.userRepo.Create(user)
	return nil
}

func (s *userService) Edit(*domain.User) error {
	return nil
}
