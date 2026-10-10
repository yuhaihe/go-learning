package handler

import (
	"net/http"

	"go-learning/projects/notes-api/internal/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Health(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			response.Fail(c, http.StatusServiceUnavailable, 503, "database unavailable")
			return
		}
		if err := sqlDB.PingContext(c.Request.Context()); err != nil {
			response.Fail(c, http.StatusServiceUnavailable, 503, "database unavailable")
			return
		}
		response.OK(c, gin.H{"status": "up"})
	}
}
