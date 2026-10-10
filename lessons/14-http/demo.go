//go:build ignore

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
	})
	fmt.Println("open http://127.0.0.1:3014/users/7")
	_ = r.Run("127.0.0.1:3014")
}
