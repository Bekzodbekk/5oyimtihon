package redis

import (
	"api-gateway/internal/pkg/load"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(conf load.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", conf.Redis.RedisHost, conf.Redis.RedisPort),
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
