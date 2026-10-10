package repo

import (
	"context"
	"errors"

	"go-learning/projects/notes-api/internal/model"

	"gorm.io/gorm"
)

type NoteRepo struct {
	db *gorm.DB
}

func NewNoteRepo(db *gorm.DB) *NoteRepo {
	return &NoteRepo{db: db}
}

func (r *NoteRepo) Create(ctx context.Context, note *model.Note) error {
	return r.db.WithContext(ctx).Create(note).Error
}

func (r *NoteRepo) GetByID(ctx context.Context, id uint) (model.Note, error) {
	var note model.Note
	err := r.db.WithContext(ctx).First(&note, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Note{}, ErrNotFound
	}
	return note, err
}

func (r *NoteRepo) List(ctx context.Context, userID uint, offset, limit int) ([]model.Note, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.Note{}).Where("user_id = ?", userID)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var notes []model.Note
	err := q.Order("id desc").Offset(offset).Limit(limit).Find(&notes).Error
	return notes, total, err
}

func (r *NoteRepo) Delete(ctx context.Context, userID, id uint) error {
	res := r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&model.Note{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
