package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var ctx = context.Background()

type redisClient struct {
	rdb redis.Client
}

func InitRedisClient() *redisClient {
	DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return &redisClient{
		rdb: *redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       DB,
		}),
	}
}

//
//var ctx = context.Background()
//
//func ExampleClient() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
