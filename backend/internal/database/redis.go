package database

import (
    "github.com/go-redis/redis/v8"
    "context"
)

func NewRedis(addr, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })
    return rdb
}
