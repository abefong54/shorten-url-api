package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	// var options *redis.Options

	// redisAddress := os.Getenv("REDIS_URL")
	// options = &redis.Options{
	// 	Addr:     redisAddress,
	// 	Password: os.Getenv("REDISPASSWORD"),
	// 	DB:       dbNo,
	// }

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDISPASSWORD"),
		DB:       dbNo,
	})
	return rdb
}
