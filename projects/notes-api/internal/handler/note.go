package handler

import (
	"errors"
	"net/http"
	"strconv"

	"go-learning/projects/notes-api/internal/middleware"
	"go-learning/projects/notes-api/internal/response"
	"go-learning/projects/notes-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Note struct {
	notes *service.NoteService
}

func NewNote(notes *service.NoteService) *Note {
	return &Note{notes: notes}
}

type createNoteReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

func (h *Note) Create(c *gin.Context) {
	var req createNoteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, 400, "invalid json")
		return
	}
	note, err := h.notes.Create(c.Request.Context(), middleware.UserID(c), req.Title, req.Content)
	if errors.Is(err, service.ErrInvalidInput) {
		response.Fail(c, http.StatusBadRequest, 400, "invalid title or content")
		return
	}
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "create failed")
		return
	}
	response.Created(c, note)
}

func (h *Note) Get(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	note, err := h.notes.Get(c.Request.Context(), middleware.UserID(c), id)
	if service.IsNotFound(err) {
		response.Fail(c, http.StatusNotFound, 404, "note not found")
		return
	}
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "get failed")
		return
	}
	response.OK(c, note)
}

func (h *Note) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	result, err := h.notes.List(c.Request.Context(), middleware.UserID(c), page, size)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "list failed")
		return
	}
	response.OK(c, result)
}

func (h *Note) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	err := h.notes.Delete(c.Request.Context(), middleware.UserID(c), id)
	if service.IsNotFound(err) {
		response.Fail(c, http.StatusNotFound, 404, "note not found")
		return
	}
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, 500, "delete failed")
		return
	}
	response.OK(c, nil)
}

func parseID(c *gin.Context) (uint, bool) {
	n, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || n == 0 {
		response.Fail(c, http.StatusBadRequest, 400, "invalid id")
		return 0, false
	}
	return uint(n), true
}
