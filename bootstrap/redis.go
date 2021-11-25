package bootstrap

import (
	"data-pipeline/config"

	"github.com/go-redis/redis"
)

func (bs *Bootstrap) initRedis() *Bootstrap {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "",
		DB:       0,
	})

	bs.Redis = client

	return bs
}
