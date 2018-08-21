package infrastructure

import (
	"context"
	"database/sql"
	"log"

	"github.com/mafuyuk/ddd-go-api-template/domain"
	"github.com/mafuyuk/ddd-go-api-template/domain/repository"
	"github.com/mafuyuk/ddd-go-api-template/infrastructure/db"

	sq "gopkg.in/Masterminds/squirrel.v1"
)

func NewUserRepository(dbmClient *db.Client, dbsClient *db.Client) repository.UserRepository {
	return &userRepository{
		dbm: dbmClient,
		dbs: dbsClient,
	}
}

type userRepository struct {
	dbm *db.Client
	dbs *db.Client
}

func (r *userRepository) WithTransaction(ctx context.Context, txFunc func(*sql.Tx) error) error {
	tx, err := r.dbm.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}

func (r *userRepository) Get(domain.UserID) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) Create(user *domain.User) error {
	log.Println("called infrastructure Create")
	// Exec Create
	_, err := sq.Insert("users").
		Columns("id", "name", "created_at", "updated_at").
		Values(user.ID, user.Name, user.CreatedAt, user.UpdatedAt).
		RunWith(r.dbm.DB).
		Exec()
	return err
}

func (r *userRepository) Update(*domain.User) error {
	return nil
}
