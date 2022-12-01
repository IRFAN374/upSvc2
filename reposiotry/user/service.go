package user

import (
	"context"

	"github.com/IRFAN374/upSvc2/models"
)

type Repository interface {
	Add(ctx context.Context, userId string, user models.User) (err error)
	Get(ctx context.Context, userId string) (user models.User, err error)
	Update(ctx context.Context, userId string) (user models.User, err error)
	Delete(ctx context.Context, userId string) (err error)
	IsExist(ctx context.Context, userId string) (ok bool, err error)
}
