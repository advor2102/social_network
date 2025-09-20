package contracts

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheI interface {
	Set(rdb *redis.Client, ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string, response interface{}) error
}
