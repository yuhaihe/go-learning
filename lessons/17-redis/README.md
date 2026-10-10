# 17. Redis 缓存

用 Redis 记下查过的数据，并设置过期时间。读的时候先 `GET`，没有再查数据库，然后 `SET`。

先确认本机 Redis 已启动，再：

```bash
export REDIS_ADDR=127.0.0.1:6379
go run lessons/17-redis/demo.go
go test ./lessons/17-redis
```

没连上 Redis 时测试会跳过，不算失败。

## 怎么写

```go
err := rdb.Set(ctx, key, value, ttl).Err()
return rdb.Get(ctx, key).Result()
```

`Set` 的最后一个参数就是过期时间。

## 练习

`Remember` 把值写入 key，带上过期时间，再读出来返回。
