package cachex

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRemember(t *testing.T) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		t.Skip("Redis 不可用，跳过：", err)
	}
	key := "lesson17:test"
	t.Cleanup(func() { rdb.Del(context.Background(), key) })

	got, err := Remember(ctx, rdb, key, "ok", time.Minute)
	if err != nil || got != "ok" {
		t.Fatalf("Remember=%q %v", got, err)
	}
}
