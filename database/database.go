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
	fmt.Println("TESTING TB")
	var options *redis.Options

	if local == "true" {
		redisAddress := os.Getenv("LOCAL_REDIS_URL")
		// rdb := redis.NewClient(&redis.Options{
		// 	Addr:     os.Getenv("LOCAL_REDIS_URL"),
		// 	Password: os.Getenv("LOCAL_REDISPASSWORD"),
		// 	DB:       dbNo,
		// })
		options = &redis.Options{
			Addr:     redisAddress,
			Password: os.Getenv("LOCAL_REDISPASSWORD"),
			DB:       dbNo,
		}

	} else {
		fmt.Println("not local")
		buildOpts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
		options = buildOpts
	}

	rdb := redis.NewClient(options)
	fmt.Println("rdb")
	fmt.Println(rdb)
	fmt.Println("____")
	return rdb
}
