# 12. context

`context.Context` 用来传递取消和超时。HTTP、数据库、Redis 的调用都应带上请求的 ctx，客户端断开时下游也停。

函数的第一个参数习惯是 `ctx context.Context`。不要把 ctx 存在结构体里。

## 怎么写

`select` 同时等两件事，谁先到就走谁：

```go
timer := time.NewTimer(d)
defer timer.Stop()
select {
case <-ctx.Done():
    return ctx.Err()
case <-timer.C:
    return nil
}
```

`ctx.Done()` 在取消或超时时关闭。`timer.C` 在 `d` 结束后收到信号。

## 命令

```bash
go run lessons/12-context/demo.go
go test ./lessons/12-context
```

## 练习

`Wait` 在 ctx 取消时返回 `ctx.Err()`，否则等到 `d` 结束返回 nil。
