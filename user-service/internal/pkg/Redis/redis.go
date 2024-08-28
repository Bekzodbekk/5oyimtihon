package redis

import (
	"context"
	"fmt"
	"user-service/internal/pkg/load"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg load.Config) (*redis.Client, error) {
	target := fmt.Sprintf("%s:%d", cfg.Redis.RedisHost, cfg.Redis.RedisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr: target,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
