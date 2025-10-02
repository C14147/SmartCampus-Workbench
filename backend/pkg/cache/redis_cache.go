package cache

import (
    "context"
    "encoding/json"
    "time"
    "github.com/go-redis/redis/v8"
)

type RedisCache struct {
    Client *redis.Client
    Ctx    context.Context
}

func NewRedisCache(client *redis.Client) *RedisCache {
    return &RedisCache{
        Client: client,
        Ctx:    context.Background(),
    }
}

func (r *RedisCache) Get(key string, dest interface{}) error {
    val, err := r.Client.Get(r.Ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), dest)
}

func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return r.Client.Set(r.Ctx, key, data, expiration).Err()
}

func (r *RedisCache) Delete(key string) error {
    return r.Client.Del(r.Ctx, key).Err()
}
