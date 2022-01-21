package db

import (
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client 

func LoadReids(c [3]string) {

	dsn := c[0]
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}