# 12. context

一次请求常常要同时等数据库、下游接口。用户关掉页面，或者时间到了，这些等待都该停。`context` 就是这份「还要不要继续」的通知，习惯命名为 `ctx`。

函数第一个参数写成 `ctx context.Context`。不要把 `ctx` 存在结构体里，它只跟这一次调用走。

## demo 里的词

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
defer cancel()
<-ctx.Done()
fmt.Println(ctx.Err())
```

`context.Background()` 是一份空白的根通知，表示「先一直做下去」。

`WithTimeout` 从它派生出一份带时限的 `ctx`。第二个参数是多久后自动取消。这里的 `time.Millisecond` 是 1 毫秒（千分之一秒）。常见的还有 `time.Second`。时限到了，`ctx` 自己变成已取消。它同时返回 `cancel`，也可以提前手动取消，不必干等。

`defer cancel()`：`defer` 把这句话推迟到当前函数结束前再执行。`WithTimeout` 内部有个计时器，调用 `cancel()` 是为了拆掉它。即使已经超时，结束前也要调用一次。

`<-ctx.Done()` 停在这里，直到 `ctx` 被取消。`Done()` 返回一个 channel，取消或超时后它会关闭，`<-` 才继续往下走。这和上一课「没人发送就一直等」是同一类等待。

`ctx.Err()` 是取消原因。demo 是超时，打印 `context deadline exceeded`。有人调用了 `cancel()` 时，原因是 `context.Canceled`。

## 怎么写

练习要同时等两件事：`ctx` 先取消，就返回 `ctx.Err()`；时间 `d` 先到，就返回 `nil`。

```go
timer := time.NewTimer(d) // d 之后，timer.C 会收到一个信号
defer timer.Stop()         // 函数结束前关掉计时器
select {
case <-ctx.Done():
    return ctx.Err()
case <-timer.C:
    return nil
}
```

`select` 谁先就绪就走谁。测试里一种情况是先 `cancel()`，应得到 `context.Canceled`；另一种是 `ctx` 不会取消，只等 1 毫秒，应得到 `nil`。

## 命令

```bash
go run lessons/12-context/demo.go
go test ./lessons/12-context
```

## 练习

`Wait` 在 `ctx` 取消时返回 `ctx.Err()`，否则等到 `d` 结束返回 `nil`。
