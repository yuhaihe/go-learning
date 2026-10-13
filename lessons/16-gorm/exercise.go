package store

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

type Note struct {
	ID     uint
	UserID uint
	Title  string
}

// CreateUserAndNote 在一个事务里创建用户和笔记。
// title 为空时返回错误，并且用户也不能留下。
func CreateUserAndNote(db *gorm.DB, name, title string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		user := User{Name: name}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if title == "" {
			return errors.New("title is empty")
		}

		return tx.Create(&Note{UserID: user.ID, Title: title}).Error
	})
}
