# 17. Redis 缓存

Redis 是放在内存里的键值库，用来记下查过的数据，并设置过期时间。读的时候先 `GET`，没有再查数据库，然后 `SET`。

## 这一课的词

`SET key value 过期时间` 写入。过期时间也叫 ttl，到点这个键就没了。`GET key` 读取；键不存在会返回错误。

`rdb.Set(ctx, key, value, ttl).Err()` 执行写入并取出错误。`rdb.Get(ctx, key).Result()` 取出值和错误。`ctx` 是上一课的取消通知，Redis 命令也要带上。

`REDIS_ADDR` 是 Redis 地址，本机一般是 `127.0.0.1:6379`。先确认本机 Redis 已启动。没连上时测试会跳过，不算失败。

```bash
export REDIS_ADDR=127.0.0.1:6379
go run lessons/17-redis/demo.go
go test ./lessons/17-redis
```

## 怎么写

```go
err := rdb.Set(ctx, key, value, ttl).Err()
return rdb.Get(ctx, key).Result()
```

`Set` 的最后一个参数就是过期时间。

## 练习

`Remember` 把值写入 key，带上过期时间，再读出来返回。
