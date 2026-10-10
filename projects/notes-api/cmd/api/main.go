package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-learning/projects/notes-api/internal/config"
	"go-learning/projects/notes-api/internal/handler"
	"go-learning/projects/notes-api/internal/httpserver"
	"go-learning/projects/notes-api/internal/model"
	"go-learning/projects/notes-api/internal/queue"
	"go-learning/projects/notes-api/internal/repo"
	"go-learning/projects/notes-api/internal/service"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if cfg.MySQLDSN == "" || cfg.JWTSecret == "" {
		log.Error("MYSQL_DSN and JWT_SECRET are required, copy .env.example to .env")
		os.Exit(1)
	}

	db, err := gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
	if err != nil {
		log.Error("open mysql", "err", err)
		os.Exit(1)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("sql db", "err", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Note{}); err != nil {
		log.Error("migrate", "err", err)
		os.Exit(1)
	}

	rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisAddr})

	authSvc := service.NewAuth(repo.NewUserRepo(db), cfg.JWTSecret)
	noteSvc := service.NewNoteService(repo.NewNoteRepo(db), repo.NewNoteCache(rdb), queue.NewAsynqEnqueuer(asynqClient), log)
	engine := httpserver.New(log, db, cfg.JWTSecret, handler.NewAuth(authSvc), handler.NewNote(noteSvc))

	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: engine}
	go func() {
		log.Info("api listening", "addr", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen", "err", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("shutdown", "err", err)
	}
	_ = asynqClient.Close()
	_ = rdb.Close()
	_ = sqlDB.Close()
	log.Info("api stopped")
}
