# 20. 配置与日志

配置从环境变量读。缺省值只用于本机开发，密码和密钥不要写进代码。

用标准库 `log/slog` 打结构化日志。请求日志至少包含方法、路径和状态码。

## 这一课的词

环境变量是进程启动时从外面传进来的配置。`os.Getenv("HTTP_ADDR")` 按名字读取，没设置时得到空字符串 `""`。地址为空就用缺省值 `:3000`。

`MYSQL_DSN` 是连 MySQL 的那一串地址，里面有账号密码，所以不要写进代码。

结构化日志不是一整句文字，而是成对的键和值。`log.Info("request", "method", method, "path", path, "status", status)` 里，第一个是消息，后面两两一组。这样以后可以按 `status` 检索，而不是从一句话里抠。

## 怎么写

```go
addr := os.Getenv("HTTP_ADDR")
if addr == "" {
    addr = ":3000"
}
log.Info("request", "method", method, "path", path, "status", status)
```

`Info` 后面是成对的键和值，不是 `Printf`。

## 命令

```bash
HTTP_ADDR=:3000 MYSQL_DSN='user:pass@tcp(127.0.0.1:3306)/notes' go run lessons/20-config/demo.go
go test ./lessons/20-config
```

## 练习

`Load` 读取 `HTTP_ADDR` 和 `MYSQL_DSN`。地址为空时用 `:3000`。

`LogRequest` 打一条包含 method、path、status 的 info 日志。
