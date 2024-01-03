package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	local := os.Getenv("LOCAL")

	if local == "true" {
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("LOCAL_REDIS_URL"),
			Password: os.Getenv("LOCAL_REDISPASSWORD"),
			DB:       dbNo,
		})

		return rdb

	} else {
		fmt.Println("not local")
		rdb := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDISPASSWORD"),
			DB:       dbNo,
		})
		return rdb
	}
}
