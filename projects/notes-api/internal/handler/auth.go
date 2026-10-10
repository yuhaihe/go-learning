package handler

import (
	"errors"
	"net/http"

	"go-learning/projects/notes-api/internal/response"
	"go-learning/projects/notes-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	auth *service.Auth
}

func NewAuth(auth *service.Auth) *Auth {
	return &Auth{auth: auth}
}

type accountReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Auth) Register(c *gin.Context) {
	var req accountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	err := h.auth.Register(c.Request.Context(), req.Username, req.Password)
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		response.Fail(c, http.StatusBadRequest, 400, "username required and password at least 6 chars")
	case errors.Is(err, service.ErrConflict):
		response.Fail(c, http.StatusConflict, 409, "username already exists")
	case err != nil:
		response.Fail(c, http.StatusInternalServerError, 500, "register failed")
	default:
		response.Created(c, gin.H{"username": req.Username})
	}
}

func (h *Auth) Login(c *gin.Context) {
	var req accountReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	token, err := h.auth.Login(c.Request.Context(), req.Username, req.Password)
	switch {
	case errors.Is(err, service.ErrUnauthorized):
		response.Fail(c, http.StatusUnauthorized, 401, "username or password incorrect")
	case err != nil:
		response.Fail(c, http.StatusInternalServerError, 500, "login failed")
	default:
		response.OK(c, gin.H{"token": token})
	}
}
