package dao

import (
	"github.com/go-redis/redis"
	"go_shop/conf"
)

var Client *redis.Client

func RedisInit() {
	conf.Conf()
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic("Redis连接失败")
	}
}
