package service

import (
	"context"
	"testing"

	"go-learning/projects/notes-api/internal/model"
	"go-learning/projects/notes-api/internal/repo"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type memCache struct {
	data map[uint]model.Note
}

func (m *memCache) Get(_ context.Context, id uint) (model.Note, bool, error) {
	note, ok := m.data[id]
	return note, ok, nil
}
func (m *memCache) Set(_ context.Context, note model.Note) error {
	m.data[note.ID] = note
	return nil
}
func (m *memCache) Del(_ context.Context, id uint) error {
	delete(m.data, id)
	return nil
}

type memQueue struct{ n int }

func (q *memQueue) EnqueueNoteCreated(context.Context, uint, string) error {
	q.n++
	return nil
}

func TestNoteFlow(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Note{}); err != nil {
		t.Fatal(err)
	}
	cache := &memCache{data: map[uint]model.Note{}}
	q := &memQueue{}
	svc := NewNoteService(repo.NewNoteRepo(db), cache, q, nil)
	ctx := context.Background()

	note, err := svc.Create(ctx, 1, "hello", "body")
	if err != nil {
		t.Fatal(err)
	}
	if q.n != 1 {
		t.Fatalf("应投递 1 次任务，得到 %d", q.n)
	}
	got, err := svc.Get(ctx, 1, note.ID)
	if err != nil || got.Title != "hello" {
		t.Fatalf("%+v %v", got, err)
	}
	if _, ok := cache.data[note.ID]; !ok {
		t.Fatal("读过之后应写入缓存")
	}
	page, err := svc.List(ctx, 1, 1, 10)
	if err != nil || page.Total != 1 {
		t.Fatalf("%+v %v", page, err)
	}
	if _, err := svc.Get(ctx, 2, note.ID); !IsNotFound(err) {
		t.Fatal("别人的笔记应视为不存在")
	}
	if err := svc.Delete(ctx, 1, note.ID); err != nil {
		t.Fatal(err)
	}
	if _, ok := cache.data[note.ID]; ok {
		t.Fatal("删除后应清掉缓存")
	}
}
