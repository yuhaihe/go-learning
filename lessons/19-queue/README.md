# 19. Asynq 队列

接口里不要做慢操作。把任务丢进队列，另起一个 worker 消费。Asynq 是任务队列库，用 Redis 存任务，本机不用再装 Kafka。

Kafka 和这里的差别主要是中间件，用法一样：生产者只负责投递，消费者单独进程跑。这一课的测试不连 Redis，只检查你生成的任务。

## 这一课的词

生产者是把任务放进队列的一方，通常是接口。worker 是另一个进程，从队列里取出任务再执行。

payload 是任务带上的数据。`json.Marshal` 把 Go 的值变成 JSON 字节。`map[string]any` 是键为字符串、值类型不固定的 map。

`asynq.NewTask("note:created", payload)` 只是做出一个任务，还没发送。`Enqueue` 才是放进 Redis。练习只做到 `NewTask`。

## 怎么写

```go
payload, err := json.Marshal(map[string]any{
    "id":    id,
    "title": title,
})
return asynq.NewTask(TaskNoteCreated, payload), nil
```

`TaskNoteCreated` 常量已经声明好了。

## 命令

```bash
export REDIS_ADDR=127.0.0.1:6379
go run lessons/19-queue/demo.go
go test ./lessons/19-queue
```

## 练习

`NewNoteTask` 生成类型为 `note:created` 的任务，载荷是 JSON：`{"id":1,"title":"..."}`。
