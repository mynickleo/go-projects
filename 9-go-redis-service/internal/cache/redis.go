package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr string) *RedisCache {
	redisDB := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisCache{
		client: redisDB,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) Set(key string, value string) error {
	return r.client.Set(r.ctx, key, value, 10*time.Second).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
