//go:build ignore

package main

import "fmt"

type CreateNote struct {
	Title   string `json:"title" binding:"required,min=1,max=100"`
	Content string `json:"content" binding:"max=2000"`
}

func main() {
	fmt.Printf("%+v\n", CreateNote{Title: "hello"})
}
