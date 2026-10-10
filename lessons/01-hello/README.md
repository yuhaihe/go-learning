# 1. 模块与命令

一个 Go 模块由 `go.mod` 声明。可运行的程序是 `package main`，并且有 `func main()`。

和 JS 不同：没有独立的「脚本运行时」。`go run` 会先编译再执行。

## 命令

```bash
go run lessons/01-hello/demo.go
go test ./lessons/01-hello
```

`demo.go` 第一行是 `//go:build ignore`，这样它不会和练习文件抢同一个包。

## 练习

实现 `Hello`。`name` 为空时返回 `Hello, World`，否则返回 `Hello, ` 加上名字。

在 `exercise.go` 里补上实现，再跑测试。`demo.go` 可以对照。
