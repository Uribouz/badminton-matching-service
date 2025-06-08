package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)


type Cache struct {
    rdb *redis.Client
}

func NewClient(host string) *Cache {
    return &Cache{rdb: redis.NewClient(&redis.Options{
        Addr:     host,
        Password: "", // no password set
        DB:       0,  // use default DB
    })}
}

func (c *Cache) SetValue(ctx context.Context, key string, value string) error {
    err := c.rdb.Set(ctx, key, value, 0).Err()
    fmt.Printf("[Set] %v: %v" + key, value)
    return err
}

func (c *Cache) GetValue (ctx context.Context, key string) (string,error) {
    value, err := c.rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key2 does not exist")
    } else if err != nil {
        return "", err
    }
    fmt.Printf("[Get] %v: %v" + key, value)
    return value, nil
}