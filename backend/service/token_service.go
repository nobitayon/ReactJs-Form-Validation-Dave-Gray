package service

import (
	"context"
	"log"

	"be-react-login/handler/model"
	"be-react-login/handler/model/apperrors"
)

type tokenService struct {
	Secret           string
	IDExpirationSecs int64
}

type TSConfig struct {
	Secret           string
	IDExpirationSecs int64
}

func NewTokenService(c *TSConfig) model.TokenService {
	return &tokenService{
		Secret:           c.Secret,
		IDExpirationSecs: c.IDExpirationSecs,
	}
}

func (s *tokenService) NewTokenFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	idToken, err := generateIDToken(u, s.Secret, s.IDExpirationSecs)
	if err != nil {
		log.Printf("error generating idToken for uid: %v. Error: %v", u.UID, err.Error())
		return nil, apperrors.NewInternal()
	}

	return &model.TokenPair{
		IDToken: idToken,
	}, nil
}
