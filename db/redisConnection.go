package db

import (
	redis "github.com/go-redis/redis"
)


func ConnectRedis(redisAddr string) (client *redis.Client, err error) {
	if len(redisAddr) < 1 {
		redisAddr = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	_, err = client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client, nil
}
