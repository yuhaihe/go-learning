# 14. Gin 与调用下游 HTTP

Gin 负责路由和 JSON。路径参数用 `c.Param`，状态码用 `http.StatusOK` 这类常量。

公司里的接口服务经常不查自己的表，而是按配置里的 host 和 timeout 去调别的 HTTP 服务。调用时带上 `context`，超时或客户端断开就停。

## 命令

```bash
go run lessons/14-http/demo.go
go test ./lessons/14-http
```

demo 会在 `:3014` 起一个临时服务。练习不占用端口，用 `httptest` 测。

## 练习

`NewRouter` 注册 `GET /users/:id`，返回 JSON `{"id":"..."}`。

`GetText` 用给定的 timeout 请求 URL，只接受 200，返回响应正文。ctx 已取消或超时要返回错误。
