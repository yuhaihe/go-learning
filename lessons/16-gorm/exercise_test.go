package store

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func openDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&User{}, &Note{}); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestCreateUserAndNote(t *testing.T) {
	db := openDB(t)
	if err := CreateUserAndNote(db, "ann", "hello"); err != nil {
		t.Fatal(err)
	}
	var n int64
	db.Model(&Note{}).Count(&n)
	if n != 1 {
		t.Fatalf("笔记数量 %d", n)
	}
}

func TestCreateUserAndNoteRollback(t *testing.T) {
	db := openDB(t)
	if err := CreateUserAndNote(db, "ann", ""); err == nil {
		t.Fatal("空标题应失败")
	}
	var n int64
	db.Model(&User{}).Count(&n)
	if n != 0 {
		t.Fatalf("事务应回滚，用户还在，数量 %d", n)
	}
}
