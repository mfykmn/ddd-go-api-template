package application

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/mafuyuk/ddd-go-api-template/domain"
	"github.com/mafuyuk/ddd-go-api-template/domain/repository"

	"github.com/rs/xid"
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

	// ID生成
	user.ID = generateUserID()

	// 時刻の埋め込み
	now := int(time.Now().Unix())
	user.CreatedAt = now
	user.UpdatedAt = now

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

func generateUserID() domain.UserID {
	id := xid.New().String()
	return domain.UserID(id)
}
