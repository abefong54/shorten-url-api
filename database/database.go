package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	var options *redis.Options

	if os.Getenv("LOCAL") == "true" {

		redisAddress := os.Getenv("REDIS_URL")
		options = &redis.Options{
			Addr:     redisAddress,
			Password: os.Getenv("REDISPASSWORD"),
			DB:       dbNo,
		}

	} else {
		buildOpts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
		options = buildOpts
	}

	rdb := redis.NewClient(options)
	return rdb
}
