# 8. 结构体、方法、组合

结构体把相关字段放在一起。方法是带接收者的函数。值接收者不能改原对象；要改字段用指针接收者。

Go 没有继承。把另一个结构体不写字段名地嵌进去，叫组合，外层可以直接用内层的字段。

## 命令

```bash
go run lessons/08-struct/demo.go
go test ./lessons/08-struct
```

## 练习

`Rect.Area` 返回宽乘高。`User` 组合 `Address`，`Label` 返回 `名字@城市`。
