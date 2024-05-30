package service

import (
	"context"
	"log"

	"be-react-login/handler/model"

	"github.com/google/uuid"
)

type userService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository: c.UserRepository,
	}
}

func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
	pw, err := hashPassword(u.Password)
	if err != nil {
		log.Printf("unable to signup for username: %v\n", u.Username)
	}

	// unnatural to mutate password here
	// better the function signature (ctx, email, password)
	// create struct user
	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}
	return nil
}
