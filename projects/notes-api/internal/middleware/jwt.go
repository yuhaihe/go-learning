package middleware

import (
	"net/http"
	"strings"

	"go-learning/projects/notes-api/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const ContextUserID = "uid"

func JWT(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		raw, ok := strings.CutPrefix(header, "Bearer ")
		if !ok || raw == "" {
			response.Fail(c, http.StatusUnauthorized, 401, "missing token")
			c.Abort()
			return
		}
		parsed, err := jwt.Parse(raw, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !parsed.Valid {
			response.Fail(c, http.StatusUnauthorized, 401, "invalid token")
			c.Abort()
			return
		}
		claims, ok := parsed.Claims.(jwt.MapClaims)
		if !ok {
			response.Fail(c, http.StatusUnauthorized, 401, "invalid token")
			c.Abort()
			return
		}
		uid, ok := claims["uid"].(float64)
		if !ok || uid == 0 {
			response.Fail(c, http.StatusUnauthorized, 401, "invalid token")
			c.Abort()
			return
		}
		c.Set(ContextUserID, uint(uid))
		c.Next()
	}
}

func UserID(c *gin.Context) uint {
	v, _ := c.Get(ContextUserID)
	id, _ := v.(uint)
	return id
}
