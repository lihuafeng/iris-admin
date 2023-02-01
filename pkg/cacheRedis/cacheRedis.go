package cacheRedis

import (
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

	return redisClient
}