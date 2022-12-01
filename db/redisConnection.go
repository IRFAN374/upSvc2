package db

import (
	"os"

	redis "github.com/go-redis/redis/v8"
)

var client *redis.Client

func ConnectRedis() *redis.Client {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
