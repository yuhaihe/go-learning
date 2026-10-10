//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hibiken/asynq"
)

func main() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	payload, _ := json.Marshal(map[string]any{"id": 1, "title": "hello"})
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
	defer client.Close()
	info, err := client.Enqueue(asynq.NewTask("note:created", payload))
	if err != nil {
		panic(err)
	}
	fmt.Println("enqueued", info.ID)
}
