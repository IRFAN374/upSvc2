package token

import (
	"context"

	"github.com/IRFAN374/upSvc2/models"
	"github.com/IRFAN374/upSvc2/reposiotry/token"
)

type Service interface {
	CreateToken(ctx context.Context, userid uint64, username string) (token models.TokenResponse, err error)
	VerifyToken(ctx context.Context, token string) (ok bool, err error)
}

type service struct {
	tokenRepo token.Repository
}

func NewService(tokenRepo token.Repository) *service {
	return &service{
		tokenRepo: tokenRepo,
	}
}

func (svc *service) CreateToken(ctx context.Context, userid uint64, username string) (token models.TokenResponse, err error) {
	return
}

func (svc *service) VerifyToken(ctx context.Context, token string) (ok bool, err error) {
	return
}
