# 11. goroutine、channel、select

`go f()` 另开一个 goroutine，和当前函数同时往下跑。goroutine 之间不要靠改同一个变量传结果，并发写会数据竞争。channel 就是这条传送带：一边 `ch <- 值` 送出去，另一边 `值 := <-ch` 接过来。

```go
ch := make(chan int)    // 无缓冲：发送方会停住，直到有人接收
ch := make(chan int, 2) // 有缓冲：缓冲区没满时，发送可以先走，值先停在里面
```

无缓冲时，发送和接收互相等，对上了才继续。有缓冲时，空位还在就可以先发送。缓冲区空了再接收，接收方会停住，除非用 `select` 的 `default` 表示「现在取不到就先干别的」。demo 里先放进 `"ping"` 会打印 `ping`；空 channel 走 `default`，打印 `empty`。

练习里 channel 的用法就是把结果送回主 goroutine：两个 goroutine 各算一半，`ch <-` 送总和，主 goroutine `<-ch` 收两次再相加。

`select` 同时等多个 channel，哪个先就绪就走哪个。不要用共享变量裸奔，需要共享时再加锁。

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
