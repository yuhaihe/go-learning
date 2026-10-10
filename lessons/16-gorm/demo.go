//go:build ignore

package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	db.Create(&User{Name: "ann"})
	var user User
	db.First(&user, 1)
	fmt.Println(user.ID, user.Name)
}
