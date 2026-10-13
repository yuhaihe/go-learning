package cachex

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Remember 写入 key，带上过期时间，再读出来返回。
func Remember(ctx context.Context, rdb *redis.Client, key, value string, ttl time.Duration) (string, error) {
	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return "", err
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
