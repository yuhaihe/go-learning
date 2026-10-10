# Go 常用命令

在仓库根目录执行。查看子命令：`go help <command>`。

## 版本与环境

| 命令 | 说明 |
|------|------|
| `go version` | 查看 Go 版本 |
| `go env GOPROXY` | 查看模块代理 |
| `go env -w GOPROXY=https://goproxy.cn,direct` | 写入 Go 自己的配置文件 |

当前 shell 里的 `export GOPROXY=...` 优先于 `go env -w`。

## 本仓库

```bash
go test ./lessons/...
go run lessons/01-hello/demo.go
go fmt ./...

# 实战。先复制 .env.example 为 .env 并填好数据库
go run ./projects/notes-api/cmd/api
go run ./projects/notes-api/cmd/worker
air
```

## 运行、编译、依赖

| 命令 | 说明 |
|------|------|
| `go run <文件或包>` | 编译并运行 |
| `go test ./...` | 运行测试 |
| `go test -run TestName ./lessons/05-slice` | 只跑一个测试 |
| `go fmt ./...` | 格式化 |
| `go mod tidy` | 按 import 整理依赖 |

## 常见问题

| 现象 | 处理 |
|------|------|
| `proxy.golang.org` 超时 | 确认 `go env GOPROXY` 是 `https://goproxy.cn,direct` |
| `go -v` 报错 | 用 `go version` |
| 练习测试失败 | 看失败信息里的「请实现」，改 `exercise.go` 后再测 |
