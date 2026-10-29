# 自己实现笔记 API

`projects/notes-api` 是做完的对照。不要先翻那里的函数体。按下面的顺序在本目录写，写完一步再打开 `notes-api` 里的同名文件看。

模块路径仍用 `go-learning/...`，不用新开 `go.mod`。先做完第 14 到第 20 课，这一份用的就是那些零件。

## 最后要长成什么样

浏览器打开首页，能注册、登录、创建笔记、翻页、查看、删除。

数据在 MySQL。按 id 查看时先问 Redis。创建成功后往队列丢一个任务，另一个进程把它打到日志里。

```text
projects/notes-mine/
  cmd/api/          进程入口，只做组装
  cmd/worker/       消费队列
  internal/handler/ 解析 HTTP，返回 JSON
  internal/middleware/ 登录校验
  internal/httpserver/ 路由
  internal/service/ 业务规则
  internal/repo/    MySQL、Redis
  internal/config/
  internal/model/
  internal/response/
  internal/queue/
  migrations/
  web/
```

请求先进 handler，业务判断在 service，MySQL 和 Redis 只出现在 repo。要登录的路由放进 JWT 分组。发日志这类慢操作丢进队列，由 `cmd/worker` 消费，不要写在 handler 里。`main` 里不要写业务判断。

仓库根目录的 `.gitignore` 已经忽略 `.env` 和 `tmp/`。`.env.example` 只留变量名和假密码，真实值放本机 `.env`。

## 第 0 步：环境和表

MySQL 里执行 `CREATE DATABASE notes;`。把仓库根目录的 `.env.example` 复制成 `.env`，填上：

- `MYSQL_DSN`：连到 `notes` 库，带 `parseTime=True`
- `REDIS_ADDR`：一般是 `127.0.0.1:6379`
- `JWT_SECRET`：随便一长串，不要提交 `.env`

建 `migrations/001_init.sql`，两张表：

- `users`：`id`、`username`（唯一）、`password_hash`、`created_at`
- `notes`：`id`、`user_id`（普通索引）、`title`、`content`、`created_at`

以后改表只追加 `002_xxx.sql`，不改已经写好的文件。

## 第 1 步：配置

文件：`internal/config/config.go`

写一个 `Config` 结构体，字段是 `HTTPAddr`、`MySQLDSN`、`RedisAddr`、`JWTSecret`。`Load()` 从环境变量读。`HTTP_ADDR` 空着时用 `:3000`，`REDIS_ADDR` 空着时用 `127.0.0.1:6379`。数据库地址和 JWT 密钥没有默认值，空了就让 `main` 退出。

这一步还没有网络。在测试里 `t.Setenv` 之后调用 `Load()`，看字段对不对。

在仓库根目录执行：

```bash
go test ./projects/notes-mine/internal/config
```

## 第 2 步：模型

文件：`internal/model/model.go`

两个结构体：`User`、`Note`。字段和上一步的表对应。`PasswordHash` 的 json 标签写成 `-`，避免以后接口把密码哈希吐出去。

GORM 靠字段名和标签建表。启动时会 `AutoMigrate`，SQL 文件仍然留着，给人看表长什么样。

## 第 3 步：统一返回

文件：`internal/response/response.go`

公司项目里常见的 JSON 是：

```json
{"code":0,"message":"ok","data":{}}
```

写 `OK`、`Created`、`Fail`。成功时 `code` 是 0。失败时 `code` 用 400、401、404、409、500，HTTP 状态码用同一个数字。

handler 只许调用这三个函数，不要在各处自己拼 `c.JSON`。

## 第 4 步：用户仓库

文件：`internal/repo/user.go`

`UserRepo` 里持有 `*gorm.DB`。两个方法：

- `Create`：插入用户
- `FindByUsername`：按用户名查。没有记录时返回一个包级错误 `ErrNotFound`，不要把 `gorm.ErrRecordNotFound` 漏到 service 外面

查询都带 `WithContext(ctx)`。

## 第 5 步：笔记仓库

文件：`internal/repo/note.go`

`NoteRepo` 四个方法：

- `Create`
- `GetByID`：没有记录同样返回 `ErrNotFound`
- `List(ctx, userID, offset, limit)`：只查这个用户的，按 `id` 倒序，同时 `Count` 总数
- `Delete(ctx, userID, id)`：条件里必须同时有 `id` 和 `user_id`。影响行数是 0 就返回 `ErrNotFound`

先不要写 Redis。用 sqlite 内存库写一个小测试：创建一条，列表总数是 1，删掉后再查是 `ErrNotFound`。sqlite 的打开方式和第 16 课一样。

在仓库根目录执行：

```bash
go test ./projects/notes-mine/internal/repo
```

## 第 6 步：缓存

文件：`internal/repo/cache.go`

`NoteCache` 持有 `*redis.Client`。键用 `note:<id>`，过期时间 5 分钟。值是笔记的 JSON。

- `Get`：Redis 说没有这个键时，返回「没找到」而不是错误。连接失败才返回错误
- `Set`
- `Del`

缓存挂了不能让接口挂。后面的 service 看到 `Get` 的错误就当没击中，继续查数据库。

## 第 7 步：队列

文件：`internal/queue/queue.go`

定义任务名 `note:created`，载荷是 `{"id":1,"title":"..."}`。

接口：

```go
type Enqueuer interface {
    EnqueueNoteCreated(ctx context.Context, id uint, title string) error
}
```

再用 Asynq 写一个实现。service 只依赖这个接口。测试里写一个假的，记一下被调用了几次。这就是第 13 课的假实现。

## 第 8 步：注册和登录

文件：`internal/service/auth.go`

`Auth` 依赖「能创建用户、能按用户名查找」的接口，不要直接依赖 `*gorm.DB`。再加一个 JWT 密钥。

`Register`：用户名不能空，密码至少 6 个字符，否则返回你自己定义的 `ErrInvalidInput`。密码用 `bcrypt` 哈希后再入库。用户名重复时返回 `ErrConflict`。

`Login`：用户不存在或密码不对，都返回 `ErrUnauthorized`，不要告诉调用方到底是哪一个错了。通过后用 HS256 签发 token，claims 里放 `uid` 和 24 小时后的 `exp`。

## 第 9 步：笔记业务

文件：`internal/service/note.go`

`NoteService` 拿三个依赖：笔记仓库接口、缓存接口、`Enqueuer`。再加一个 `*slog.Logger`。

`Create`：标题不能空，标题不超过 100 字，内容不超过 2000 字。先入库，再投递队列。队列失败只打日志，笔记仍然返回成功。慢操作失败不应该把已经写进数据库的请求变成 500。

`Get`：先读缓存。击中了但 `user_id` 不是当前用户，返回 `ErrNotFound`，不要把别人的笔记交出去。未击中就查库，确认是自己的，再写入缓存。

`List`：`page` 小于 1 当成 1，`size` 小于 1 当成 10，大于 50 当成 50。偏移量是 `(page-1)*size`。

`Delete`：仓库删除成功后删缓存。

用第 5 步的 sqlite 仓库，加上内存版缓存和假队列，测这条链路：创建会投递 1 次；自己能读到并且缓存里有；别的用户读同一条得到 `ErrNotFound`；删除后缓存没了。

在仓库根目录执行：

```bash
go test ./projects/notes-mine/internal/service
```

## 第 10 步：登录中间件

文件：`internal/middleware/jwt.go`

读请求头 `Authorization`。必须以 `Bearer ` 开头。用同一个密钥解析，取出 `uid`，放进 `gin.Context`。缺 token、签名不对、过期，都返回 401，并 `Abort`。

再写一个 `UserID(c)`，给 handler 取这个 id。

## 第 11 步：handler

文件：`internal/handler/auth.go`、`note.go`、`health.go`

handler 只做三件事：绑定 JSON 或路径参数、调用 service、选择返回。

先做 `GET /health`：对数据库 `Ping`，失败返回 503。这一步通了，再加登录和业务接口。

- `POST /register`、`POST /login`：绑定 `username`、`password`
- `POST /notes`：绑定 `title`、`content`，用户 id 从中间件拿
- `GET /notes/:id`、`DELETE /notes/:id`：id 不是正整数就 400
- `GET /notes`：查询参数 `page`、`size`

`ErrInvalidInput` 对应 400，`ErrUnauthorized` 对应 401，`ErrConflict` 对应 409，`ErrNotFound` 对应 404。其它错误 500，日志可以打，响应里不要带数据库原文。

## 第 12 步：路由

文件：`internal/httpserver/server.go`

`New` 接收 logger、数据库、JWT 密钥、两个 handler。路由集中在这里注册。

注册顺序：

1. Recovery 和请求日志。日志里至少有方法、路径、状态码、耗时
2. 开发用的跨域：允许 `Authorization` 和 `Content-Type`，`OPTIONS` 直接 204
3. `GET /health`、`GET /`
4. `POST /register`、`POST /login` 不套 JWT
5. `/notes` 分组套上 JWT，里面是创建、列表、详情、删除

`GET /` 先返回一句纯文本也行。页面留到第 15 步。

## 第 13 步：把进程组装起来

文件：`cmd/api/main.go`

顺序不要换：

1. `godotenv.Load()`，读到 `.env`
2. `config.Load()`，缺 `MYSQL_DSN` 或 `JWT_SECRET` 就退出
3. 建 `slog`
4. 连接 MySQL，`AutoMigrate` 用户和笔记
5. 连接 Redis，再建 Asynq 客户端
6. `repo` 交给 `service`，`service` 交给 `handler`，`handler` 交给 `httpserver.New`
7. `http.Server` 放进 goroutine 里 `ListenAndServe`
8. 等待 `SIGINT` 或 `SIGTERM`
9. 用 5 秒的 context 调 `Shutdown`，然后关掉 Asynq、Redis、数据库

`Shutdown` 会停止接收新请求，等正在处理的请求结束。先关数据库，正在用连接的请求会失败。

在仓库根目录执行：

```bash
go run ./projects/notes-mine/cmd/api
```

`curl http://127.0.0.1:3000/health` 应返回 `"status":"up"`。再用 curl 走一遍注册、登录、带 token 创建笔记。

改代码想自动重启时，把仓库根目录 `.air.toml` 里的编译目标改成你的入口，再在仓库根目录执行 `air`：

```toml
cmd = "go build -o ./tmp/main ./projects/notes-mine/cmd/api"
bin = "./tmp/main"
```

现在这份 `.air.toml` 编的是 `notes-api`。worker 不走 `air`，另开一个终端跑。

## 第 14 步：worker

文件：`cmd/worker/main.go`

另起一个进程。用同一个 `REDIS_ADDR` 建 Asynq server，注册 `note:created`。处理函数把载荷里的 id 和 title 打到日志。

```bash
go run ./projects/notes-mine/cmd/worker
```

再创建一条笔记。worker 的终端里应出现那条日志。API 进程不要在请求里同步写这条日志。

## 第 15 步：页面

文件：`web/index.html`，用 `web/embed.go` 嵌进二进制。

页面上要有：注册、登录、创建、按页刷新列表、查看一条、删除一条。登录后把 token 放进 `localStorage`，之后的 `fetch` 都带 `Authorization: Bearer ...`。

`GET /` 改为返回这段 HTML。浏览器打开 http://127.0.0.1:3000 ，不看代码，把上面的操作点一遍。

## 做完再对照

卡住时再打开 `projects/notes-api` 里同名文件，看一层怎么把依赖交到下一层。不要整文件复制。

对照时重点看这三处：

- `cmd/api/main.go` 的组装和退出顺序
- `service/note.go` 里缓存未击中、用户不匹配、队列失败分别怎么处理
- `httpserver/server.go` 里哪些路由在 JWT 分组外面
