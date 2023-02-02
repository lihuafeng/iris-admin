package cacheRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)
var (
	once sync.Once
	redisClient *redis.Client
)

func Instance() *redis.Client {
	once.Do(func() {
		Options := &redis.Options{
			Addr:	  "localhost:6379",
			Password: "", // no password set
			DB:		  0,  // use default DB
		}
		redisClient = redis.NewClient(Options)
	})
	var ctx = context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil{
		panic(err)
	}
	return redisClient
}