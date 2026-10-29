package repo

import (
	"context"
	"errors"
	"go-learning/projects/notes-mine/internal/model"

	"gorm.io/gorm"
)

// `NoteRepo` 四个方法：

// - `Create`
// - `GetByID`：没有记录同样返回 `ErrNotFound`
// - `List(ctx, userID, offset, limit)`：只查这个用户的，按 `id` 倒序，同时 `Count` 总数
// - `Delete(ctx, userID, id)`：条件里必须同时有 `id` 和 `user_id`。影响行数是 0 就返回 `ErrNotFound`

// 先不要写 Redis。用 sqlite 内存库写一个小测试：创建一条，列表总数是 1，删掉后再查是 `ErrNotFound`。sqlite 的打开方式和第 16 课一样。

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

	return note, nil
}

func (r *NoteRepo) List(ctx context.Context, userId uint, offset, limit int) ([]model.Note, int64, error) {
	q := r.db.WithContext(ctx).Model(&model.Note{}).Where("user_id=?", userId)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var notes []model.Note
	err := q.Order("id desc").Offset(offset).Limit(limit).Find(&notes).Error
	return notes, total, err
}

func (r *NoteRepo) Delete(ctx context.Context, userID uint, id uint) error {
	res := r.db.WithContext(ctx).Where("id=? AND user_id=?", id, userID).Delete(&model.Note{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
