package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-learning/projects/notes-api/internal/model"

	"github.com/redis/go-redis/v9"
)

type NoteCache struct {
	rdb *redis.Client
	ttl time.Duration
}

func NewNoteCache(rdb *redis.Client) *NoteCache {
	return &NoteCache{rdb: rdb, ttl: 5 * time.Minute}
}

func noteKey(id uint) string {
	return fmt.Sprintf("note:%d", id)
}

func (c *NoteCache) Get(ctx context.Context, id uint) (model.Note, bool, error) {
	raw, err := c.rdb.Get(ctx, noteKey(id)).Bytes()
	if errors.Is(err, redis.Nil) {
		return model.Note{}, false, nil
	}
	if err != nil {
		return model.Note{}, false, err
	}
	var note model.Note
	if err := json.Unmarshal(raw, &note); err != nil {
		return model.Note{}, false, err
	}
	return note, true, nil
}

func (c *NoteCache) Set(ctx context.Context, note model.Note) error {
	raw, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return c.rdb.Set(ctx, noteKey(note.ID), raw, c.ttl).Err()
}

func (c *NoteCache) Del(ctx context.Context, id uint) error {
	return c.rdb.Del(ctx, noteKey(id)).Err()
}
