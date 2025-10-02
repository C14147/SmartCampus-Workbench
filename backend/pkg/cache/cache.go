package cache

import (
    "time"
)

type Cache interface {
    Get(key string, dest interface{}) error
    Set(key string, value interface{}, expiration time.Duration) error
    Delete(key string) error
}
