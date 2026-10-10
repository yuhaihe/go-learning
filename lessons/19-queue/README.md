# 19. Asynq 队列

接口里不要做慢操作。把任务丢进队列，另起一个 worker 消费。Asynq 用 Redis 当队列，本机不用再装 Kafka。

Kafka 和这里的差别主要是中间件，用法一样：生产者只负责投递，消费者单独进程跑。这一课的测试不连 Redis，只检查你生成的任务。

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
