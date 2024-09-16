package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

var CtxBg = context.Background()

type redisClient struct {
	Rdb redis.Client
}

var redisInstance *redisClient

func InitRedisClient() {
	DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisInstance = &redisClient{
		Rdb: *redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       DB,
		}),
	}

	// Test redis
	//	RedisInstance.Rdb.Set(ctx, "key", "value", 0)
}

func GetRedisInstance() *redis.Client {
	return &redisInstance.Rdb
}
