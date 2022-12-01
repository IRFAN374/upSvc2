package user

import (
	"context"

	"github.com/IRFAN374/upSvc2/token"
)

type Service interface {
	Register(ctx context.Context, username string, password string) (userId string, err error) // remove userId from here
	Login(ctx context.Context, username string, password string) (err error)                   // add login response here
}

type service struct {
	tokenSvc token.Service
}

func NewService(tokenSvc token.Service) *service {
	return &service{
		tokenSvc: tokenSvc,
	}
}

func (svc *service) Register(ctx context.Context, username string, password string) (userId string, err error) {
	return
}

func (svc *service) Login(ctx context.Context, username string, password string) (err error) {
	return
}
