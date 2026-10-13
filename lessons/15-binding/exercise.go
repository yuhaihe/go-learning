package bind

import "github.com/gin-gonic/gin"

type CreateNote struct {
	Title   string `json:"title" binding:"required,min=1,max=100"`
	Content string `json:"content" binding:"max=2000"`
}

// ParseCreateNote 绑定 JSON。校验失败时返回错误。
func ParseCreateNote(c *gin.Context) (CreateNote, error) {
	var req CreateNote
	if err := c.ShouldBindJSON(&req); err != nil {
		return CreateNote{}, err
	}

	return req, nil
}
