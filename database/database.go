package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	var options *redis.Options

	if os.Getenv("LOCAL") == "true" {

		redisAddress := os.Getenv("DB_ADDRESS")
		options = &redis.Options{
			Addr:     redisAddress,
			Password: os.Getenv("DB_PASS"),
			DB:       dbNo,
		}

	} else {
		buildOpts, err := redis.ParseURL(os.Getenv("REDISCLOUD_URL"))
		if err != nil {
			panic(err)
		}
		options = buildOpts
	}

	fmt.Println("local: DB_ADDRESS: ", os.Getenv("DB_ADDRESS"))
	fmt.Println("heroku: REDISCLOUD_URL: ", os.Getenv("REDISCLOUD_URL"))

	rdb := redis.NewClient(options)
	return rdb
}
