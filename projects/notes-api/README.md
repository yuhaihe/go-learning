# 笔记 API

把前面的课收成一个小服务：注册登录、笔记增删查、Redis 缓存、创建后丢进队列。

当前目录是做完的对照。要自己写一遍，按 [实现步骤.md](实现步骤.md) 在 `projects/notes-mine` 里从第 0 步做，不要先看这里的函数体。

## 和公司项目的对照

| 这里 | `bus-smart-api` |
|------|-----------------|
| handler | controller |
| service | business |
| `response.OK` / `Fail` | `pontos.Success` / `Error` |
| 环境变量 | toml 里的 host、timeout |

他们的 `service` 目录多半是在调别的 HTTP 服务。这个项目自己存 MySQL，是为了把 GORM 也练到。

## 启动

1. 本机 MySQL 建库：`CREATE DATABASE notes;`
2. 复制仓库根目录的 `.env.example` 为 `.env`，填上 `MYSQL_DSN`、`REDIS_ADDR`、`JWT_SECRET`。
3. 也可以手动执行 [`migrations/001_init.sql`](migrations/001_init.sql)。程序启动时会再做一次 AutoMigrate。
4. 开两个终端，都在仓库根目录：

```bash
go run ./projects/notes-api/cmd/api
go run ./projects/notes-api/cmd/worker
```

浏览器打开 http://127.0.0.1:3000 。改 Go 代码可用 `air` 自动重启 API。worker 需要另开终端。

`GET /health` 能通，说明进程和数据库都在。

## 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 数据库探活 |
| POST | `/register` | 注册，密码至少 6 位 |
| POST | `/login` | 返回 JWT |
| POST | `/notes` | 创建，需要 token，并投递队列任务 |
| GET | `/notes?page=&size=` | 只看自己的，默认第 1 页、每页 10 条，最大 50 |
| GET | `/notes/:id` | 先读 Redis，没有再查库并回写，5 分钟过期 |
| DELETE | `/notes/:id` | 删除并清缓存 |

请求头：`Authorization: Bearer <token>`。

返回体：`{"code":0,"message":"ok","data":...}`。失败时 `code` 非 0，HTTP 状态码也是 4xx 或 5xx。

## 组装顺序

`cmd/api/main.go` 按 [`docs/从零建项目.md`](../../docs/从零建项目.md) 来：配置、日志、MySQL、Redis、路由、监听、收到信号后先停请求再关连接。
