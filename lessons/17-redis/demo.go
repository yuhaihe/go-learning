//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	ctx := context.Background()
	if err := rdb.Set(ctx, "lesson17", "ok", time.Minute).Err(); err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "lesson17").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
