package httpserver

import (
	"log/slog"
	"net/http"
	"time"

	"go-learning/projects/notes-api/internal/handler"
	"go-learning/projects/notes-api/internal/middleware"
	"go-learning/projects/notes-api/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(log *slog.Logger, db *gorm.DB, secret string, auth *handler.Auth, notes *handler.Note) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), accessLog(log), cors())
	r.GET("/health", handler.Health(db))
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", web.IndexHTML)
	})
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

	group := r.Group("/notes", middleware.JWT(secret))
	group.POST("", notes.Create)
	group.GET("", notes.List)
	group.GET("/:id", notes.Get)
	group.DELETE("/:id", notes.Delete)
	return r
}

func accessLog(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"cost", time.Since(start).String(),
		)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
