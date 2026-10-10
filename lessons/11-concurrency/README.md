# 11. goroutine、channel、select

`go f()` 启动一个 goroutine。channel 用来在 goroutine 之间传值。无缓冲 channel 的发送和接收会互相等待。

`select` 同时等多个 channel。不要用共享变量裸奔，需要共享时再加锁。

## 怎么写

空切片先返回 0。其余切成两半，各开一个 goroutine 求和，用 channel 把结果送回来：

```go
mid := len(nums) / 2
ch := make(chan int, 2)
go func() {
    ch <- add(nums[:mid])
}()
go func() {
    ch <- add(nums[mid:])
}()
return <-ch + <-ch
```

`add` 自己写一个普通的 for 循环即可。`nums[:mid]` 是前半段，`nums[mid:]` 是后半段。

## 命令

```bash
go run lessons/11-concurrency/demo.go
go test ./lessons/11-concurrency
```

## 练习

`Sum` 把切片切成两半，各自用 goroutine 求和，再把两个结果加起来。空切片返回 0。
