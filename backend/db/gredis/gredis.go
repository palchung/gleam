package gredis

import (
	"github.com/go-redis/redis/v8"
	"thefreepress/tool/setting"
)

func Setup() *redis.Client {
	redisClient := redis.NewClient(&redis.Options {
		Addr:		setting.RedisSetting.Host,
		Password:	setting.RedisSetting.Password,
		DB:			0,
	})
	return redisClient
}

