# 20. 配置与日志

配置从环境变量读。缺省值只用于本机开发，密码和密钥不要写进代码。

用标准库 `log/slog` 打结构化日志。请求日志至少包含方法、路径和状态码。

## 命令

```bash
HTTP_ADDR=:3000 MYSQL_DSN='user:pass@tcp(127.0.0.1:3306)/notes' go run lessons/20-config/demo.go
go test ./lessons/20-config
```

## 练习

`Load` 读取 `HTTP_ADDR` 和 `MYSQL_DSN`。地址为空时用 `:3000`。

`LogRequest` 打一条包含 method、path、status 的 info 日志。
