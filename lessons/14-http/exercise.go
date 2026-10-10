package httpx

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// UserHandler 处理 GET /users/:id，返回 JSON {"id":"..."}。
func UserHandler(c *gin.Context) {
	panic("TODO")
}

// NewRouter 注册 GET /users/:id。
func NewRouter() *gin.Engine {
	panic("TODO")
}

// GetText 在 timeout 内请求 URL。只接受 200，返回响应正文。
func GetText(ctx context.Context, rawURL string, timeout time.Duration) (string, error) {
	panic("TODO")
}
