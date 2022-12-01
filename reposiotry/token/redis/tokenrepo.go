package tokenredis

import (
	"context"
	"fmt"

	"github.com/IRFAN374/upSvc2/models"
	redis "github.com/go-redis/redis/v8"
)

const (
	TOKEN = "token"
)

type repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *repository {
	return &repository{
		client: client,
	}
}

func (repo *repository) getKey(userId string) string {
	return fmt.Sprintf("%s-%s", TOKEN, userId)
}

func (repo *repository) Add(ctx context.Context, userId string, payload models.PayloadRequest) ( err error) {

	_ = repo.getKey(userId)
	return
}

func (repo *repository) Get(ctx context.Context, userId string) (token models.TokenResponse, err error) {
	_ = repo.getKey(userId)
	return
}

func (repo *repository) Update(ctx context.Context, userId string) (token models.TokenResponse, err error) {

	_ = repo.getKey(userId)
	return
}

func (repo *repository) Delete(ctx context.Context, userId string) (err error) {

	_ = repo.getKey(userId)
	return
}

func (repo *repository) IsExist(ctx context.Context, userId string) (ok bool, err error) {

	_ = repo.getKey(userId)
	return
}
