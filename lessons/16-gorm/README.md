# 16. GORM 与事务

GORM 包在 `database/sql` 上面。模型是结构体，`Create` 插入，`First` 按主键查。

多步写入用 `Transaction`。回调返回错误时，整段回滚。

练习用 sqlite 内存库，所以不连 MySQL 也能测。真实项目用 MySQL：

```bash
export MYSQL_DSN='user:pass@tcp(127.0.0.1:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local'
```

```go
db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
```

## 命令

```bash
go run lessons/16-gorm/demo.go
go test ./lessons/16-gorm
```

## 练习

`CreateUserAndNote` 在一个事务里创建用户和笔记。`title` 为空时返回错误，并且用户也不能留下。
