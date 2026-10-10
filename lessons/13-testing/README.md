# 13. 表驱动测试与假实现

表驱动测试把输入和期望写进切片，一个循环跑完。公司里的单测大多长这样。

业务函数依赖接口，测试里写一个假的实现，不连数据库。这就是「包怎么拆」的原因：service 不直接 new 一个 MySQL 客户端。

## 命令

```bash
go test ./lessons/13-testing
```

## 练习

`Store.Name` 按 id 查名字。`Greet` 查到后返回 `hi ` 加名字。查不到就把错误原样返回。

测试里的 `fakeStore` 就是假实现，不要在练习函数里访问数据库。

## 怎么写

`Greet` 只做这一件事：

```go
name, err := s.Name(id)
if err != nil {
    return "", err
}
return "hi " + name, nil
```

表和假实现已经写在 `exercise_test.go`，这一课你只补 `Greet`。
