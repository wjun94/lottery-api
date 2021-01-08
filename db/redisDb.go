package db

import (
	"mango-api/config"
	"github.com/go-redis/redis"
)

var client *redis.Client

func RedisInit() *redis.Client {
	conf := config.Config.Redis
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     conf.Addr,
			Password: conf.Password, // no password set
			DB:       conf.Base,     // use default DB
		})
	}
	return client
}
