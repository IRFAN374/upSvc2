package usermongo

import (
	"context"

	"github.com/IRFAN374/upSvc2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoReposiotry struct {
	client *mongo.Client
	dbName string
}

func NewMongoReposiotry(mongoClient *mongo.Client, dbName string) *mongoReposiotry {
	return &mongoReposiotry{
		client: mongoClient,
		dbName: dbName,
	}
}

func (mr *mongoReposiotry) Add(ctx context.Context, userId string, user models.User) (err error) {

	return
}

func (mr *mongoReposiotry) Get(ctx context.Context, userId string) (user models.User, err error) {
	return
}

func (mr *mongoReposiotry) Update(ctx context.Context, userId string) (user models.User, err error) {
	return
}

func (mr *mongoReposiotry) Delete(ctx context.Context, userId string) (err error) {
	return
}

func (mr *mongoReposiotry) IsExist(ctx context.Context, userId string) (ok bool, err error) {
	return
}
