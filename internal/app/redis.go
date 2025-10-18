package app

import (
	s_redis "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var RedisStore s_redis.Store

func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     Conf.Redis.Addr,
		Password: "",
		DB:       0,
	})
	if _, err := RDB.Ping(RDB.Context()).Result(); err != nil {
		return err
	}

	var err error
	RedisStore, err = s_redis.NewStore(10, "tcp", Conf.Redis.Addr, "", "", []byte("secret"))
	if err != nil {
		return err
	}
	return nil
}
