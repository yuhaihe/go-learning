# 16. GORM 与事务

GORM 是操作数据库的库，你写结构体，它负责插入和查询。多步写入用 `Transaction`。回调返回错误时，整段回滚，前面插入的也不会留下。

## 这一课的词

模型就是一个结构体，字段对应表的列。`Create` 插入一行。`First` 按主键查出第一行。`AutoMigrate` 按结构体把表建出来，demo 里用它。

`db.Transaction(func(tx *gorm.DB) error { ... })` 里的 `tx` 是这一次事务。都用 `tx`，不要用外面的 `db`。函数返回 `nil` 才提交；返回 `error` 就回滚。

练习用 sqlite 内存库，连接串是 `:memory:`，数据只在这次测试里，不连 MySQL 也能测。DSN 是连数据库的那一串地址。真实项目用 MySQL：

```bash
export MYSQL_DSN='user:pass@tcp(127.0.0.1:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local'
```

`charset=utf8mb4` 是字符集。`parseTime=True` 让数据库里的时间能读成 Go 的 `time.Time`。`loc=Local` 用本机时区。

```go
db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
```

## 怎么写

`Transaction` 的回调返回 error，整段会回滚：

```go
return db.Transaction(func(tx *gorm.DB) error {
    user := User{Name: name}
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    if title == "" {
        return errors.New("empty title") // 此时用户也不会留下
    }
    return tx.Create(&Note{UserID: user.ID, Title: title}).Error
})
```

测试自己开 sqlite，你不用连 MySQL。

## 命令

```bash
go run lessons/16-gorm/demo.go
go test ./lessons/16-gorm
```

## 练习

`CreateUserAndNote` 在一个事务里创建用户和笔记。`title` 为空时返回错误，并且用户也不能留下。
