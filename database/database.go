package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	local := os.Getenv("LOCAL")
	var options *redis.Options

	if local == "true" {
		redisAddress := os.Getenv("LOCAL_REDIS_URL")
		options = &redis.Options{
			Addr:     redisAddress,
			Password: os.Getenv("LOCAL_REDISPASSWORD"),
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
