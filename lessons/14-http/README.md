# 14. Gin 与调用下游 HTTP

Gin 是 HTTP 框架，负责「哪个 URL 交给哪个函数」以及把结果写成 JSON。公司里的接口经常不查自己的表，而是按配置里的地址和超时去调别的 HTTP 服务。调用时带上 `context`，超时或客户端断开就停。

## 这一课的词

`gin.New()` 创建一个空路由。`r.GET("/users/:id", UserHandler)` 注册 GET。`:id` 是路径参数，`/users/7` 里的 `7` 用 `c.Param("id")` 取出。

`c.JSON(http.StatusOK, gin.H{"id": id})` 返回 JSON。`http.StatusOK` 就是数字 200。`gin.H` 是 `map[string]any` 的别名，用来拼一个 JSON 对象。

`httptest` 是标准库里的假 HTTP，练习用它发请求，不必真的占用端口。demo 才会在 `:3014` 起一个临时服务。

`http.NewRequestWithContext` 做出一个带着 `ctx` 的请求。`http.DefaultClient.Do` 把请求发出去。`defer resp.Body.Close()` 在函数结束前关掉响应体。`io.ReadAll` 把正文读成字节。状态码不是 200 时要返回错误。

## 命令

```bash
go run lessons/14-http/demo.go
go test ./lessons/14-http
```

demo 会在 `:3014` 起一个临时服务。练习不占用端口，用 `httptest` 测。

## 练习

`NewRouter` 注册 `GET /users/:id`，返回 JSON `{"id":"..."}`。

`GetText` 用给定的 timeout 请求 URL，只接受 200，返回响应正文。ctx 已取消或超时要返回错误。

## 怎么写

路由：

```go
r := gin.New()
r.GET("/users/:id", UserHandler)
// UserHandler 里：
c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
```

带超时的 HTTP 调用：

```go
ctx, cancel := context.WithTimeout(ctx, timeout)
defer cancel()
req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
resp, err := http.DefaultClient.Do(req)
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
```

状态码不是 200 时返回错误。ctx 被取消时 `Do` 会自己返回错误。
