package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(
	host string, port string, password string, db int,
) *RedisClient {

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) Set(key string, val any, exp time.Duration) error {
	status := r.Client.Set(context.Background(), key, val, exp)

	return status.Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	string := r.Client.Get(context.Background(), key)

	return string.Result()
}
