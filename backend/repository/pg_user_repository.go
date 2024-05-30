package repository

import (
	"context"
	"log"

	"be-react-login/handler/model"
	"be-react-login/handler/model/apperrors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type pGUserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &pGUserRepository{
		DB: db,
	}
}

func (r *pGUserRepository) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (username,password) VALUES ($1,$2) RETURNING *"

	if err := r.DB.GetContext(ctx, u, query, u.Username, u.Password); err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with username: %v. Reason: %v\n", u.Username, err.Code.Name())
			return apperrors.NewConflict("email", u.Username)
		}
		log.Printf("Could not create a user with username: %v. Reason: %v\n", u.Username, err)
		return apperrors.NewInternal()
	}
	return nil
}

func (r *pGUserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid=$1"

	if err := r.DB.GetContext(ctx, user, query, uid); err != nil {
		return user, apperrors.NewNotFound("uid", uid.String())
	}
	return user, nil
}
