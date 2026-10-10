package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"go-learning/projects/notes-api/internal/config"
	"go-learning/projects/notes-api/internal/queue"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: cfg.RedisAddr}, asynq.Config{Concurrency: 2})
	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.TypeNoteCreated, func(ctx context.Context, task *asynq.Task) error {
		var payload queue.Payload
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			return err
		}
		log.Info("note created", "id", payload.ID, "title", payload.Title)
		return nil
	})
	log.Info("worker listening", "redis", cfg.RedisAddr)
	if err := srv.Run(mux); err != nil {
		log.Error("worker", "err", err)
		os.Exit(1)
	}
}
