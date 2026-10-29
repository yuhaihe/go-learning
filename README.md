# Go 学习仓库

面向公司里常见的 Go 后端：语言、Gin、MySQL、Redis、JWT、队列，最后做一个能本地跑的笔记 API。

按课号学。不要跳。每一课先读 `README.md`，再跑 demo，再看练习并跑测试。

## 怎么学

```bash
# 第 1 课示例
go run lessons/01-hello/demo.go

# 第 1 课练习。写完前测试会失败，这是正常的
go test ./lessons/01-hello
```

`exercise.go` 只留了函数签名和类型，函数体是 `panic("TODO")`。按该课 README 补上实现，再跑 `go test`。`demo.go` 是可以对照的示例。

和 AI 一起学时，先自己写。测试失败再把报错贴过去，不要第一句就要完整答案。

实战放在全部课后。先按 [`projects/notes-mine/README.md`](projects/notes-mine/README.md) 自己写。做完的对照在 [`projects/notes-api`](projects/notes-api/README.md)，卡住再看。

## 课表

语言

1. [模块与命令](lessons/01-hello/README.md)
2. [变量与类型](lessons/02-variables/README.md)
3. [流程控制](lessons/03-flow/README.md)
4. [函数](lessons/04-functions/README.md)
5. [数组与 slice](lessons/05-slice/README.md)
6. [map](lessons/06-map/README.md)
7. [指针](lessons/07-pointer/README.md)
8. [结构体与方法](lessons/08-struct/README.md)
9. [接口](lessons/09-interface/README.md)
10. [错误](lessons/10-error/README.md)
11. [并发](lessons/11-concurrency/README.md)
12. [context](lessons/12-context/README.md)
13. [测试与假实现](lessons/13-testing/README.md)

常用库

14. [Gin 与调用下游 HTTP](lessons/14-http/README.md)
15. [参数绑定与校验](lessons/15-binding/README.md)
16. [GORM 与事务](lessons/16-gorm/README.md)
17. [Redis 缓存](lessons/17-redis/README.md)
18. [JWT](lessons/18-jwt/README.md)
19. [Asynq 队列](lessons/19-queue/README.md)
20. [配置与日志](lessons/20-config/README.md)

命令备忘见 [docs/go常用命令.md](docs/go常用命令.md)。

## 一次跑完全部练习

```bash
go test ./lessons/...
```

练习还没写时，`go test` 会因 `panic("TODO")` 失败。第 17 课连不上 Redis 时会跳过。第 16 课练习用 sqlite 内存库，不依赖本机 MySQL。
