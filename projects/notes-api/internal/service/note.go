package service

import (
	"context"
	"errors"
	"log/slog"
	"unicode/utf8"

	"go-learning/projects/notes-api/internal/model"
	"go-learning/projects/notes-api/internal/queue"
	"go-learning/projects/notes-api/internal/repo"
)

type NoteStore interface {
	Create(ctx context.Context, note *model.Note) error
	GetByID(ctx context.Context, id uint) (model.Note, error)
	List(ctx context.Context, userID uint, offset, limit int) ([]model.Note, int64, error)
	Delete(ctx context.Context, userID, id uint) error
}

type NoteCache interface {
	Get(ctx context.Context, id uint) (model.Note, bool, error)
	Set(ctx context.Context, note model.Note) error
	Del(ctx context.Context, id uint) error
}

type NoteService struct {
	notes NoteStore
	cache NoteCache
	queue queue.Enqueuer
	log   *slog.Logger
}

func NewNoteService(notes NoteStore, cache NoteCache, q queue.Enqueuer, log *slog.Logger) *NoteService {
	if log == nil {
		log = slog.Default()
	}
	return &NoteService{notes: notes, cache: cache, queue: q, log: log}
}

type Page struct {
	List  []model.Note `json:"list"`
	Total int64        `json:"total"`
	Page  int          `json:"page"`
	Size  int          `json:"size"`
}

func (s *NoteService) Create(ctx context.Context, userID uint, title, content string) (model.Note, error) {
	if title == "" || utf8.RuneCountInString(title) > 100 || utf8.RuneCountInString(content) > 2000 {
		return model.Note{}, ErrInvalidInput
	}
	note := model.Note{UserID: userID, Title: title, Content: content}
	if err := s.notes.Create(ctx, &note); err != nil {
		return model.Note{}, err
	}
	if err := s.queue.EnqueueNoteCreated(ctx, note.ID, note.Title); err != nil {
		s.log.Error("enqueue note created", "id", note.ID, "err", err)
	}
	return note, nil
}

func (s *NoteService) Get(ctx context.Context, userID, id uint) (model.Note, error) {
	if s.cache != nil {
		note, ok, err := s.cache.Get(ctx, id)
		if err != nil {
			s.log.Error("cache get", "id", id, "err", err)
		} else if ok {
			if note.UserID != userID {
				return model.Note{}, repo.ErrNotFound
			}
			return note, nil
		}
	}
	note, err := s.notes.GetByID(ctx, id)
	if err != nil {
		return model.Note{}, err
	}
	if note.UserID != userID {
		return model.Note{}, repo.ErrNotFound
	}
	if s.cache != nil {
		if err := s.cache.Set(ctx, note); err != nil {
			s.log.Error("cache set", "id", id, "err", err)
		}
	}
	return note, nil
}

func (s *NoteService) List(ctx context.Context, userID uint, page, size int) (Page, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	if size > 50 {
		size = 50
	}
	notes, total, err := s.notes.List(ctx, userID, (page-1)*size, size)
	if err != nil {
		return Page{}, err
	}
	if notes == nil {
		notes = []model.Note{}
	}
	return Page{List: notes, Total: total, Page: page, Size: size}, nil
}

func (s *NoteService) Delete(ctx context.Context, userID, id uint) error {
	if err := s.notes.Delete(ctx, userID, id); err != nil {
		return err
	}
	if s.cache != nil {
		if err := s.cache.Del(ctx, id); err != nil {
			s.log.Error("cache del", "id", id, "err", err)
		}
	}
	return nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, repo.ErrNotFound)
}
