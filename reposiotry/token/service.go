package token

import (
	"context"

	"github.com/IRFAN374/upSvc2/models"
)

type Repository interface {
	Add(ctx context.Context, userId string, payload models.PayloadRequest) (err error)
	Get(ctx context.Context, userId string) (token models.TokenResponse, err error)
	Update(ctx context.Context, userId string) (token models.TokenResponse, err error)
	Delete(ctx context.Context, userId string) (err error)
	IsExist(ctx context.Context, userId string) (ok bool, err error)
}
