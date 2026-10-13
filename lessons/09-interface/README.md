# 9. 接口

接口只列方法，不写具体怎么做。`Speaker` 只要求有 `Speak() string`。`Dog` 有这个方法，就能放进 `Speaker` 变量，不用写 `implements`。

`var s Speaker = Dog{}` 里，`s` 的类型是接口，里面装的是一只 `Dog`。`s == nil` 表示什么都没装。业务代码只认接口，测试时可以换一个假的实现。

## 怎么写

```go
func (Dog) Speak() string {
    return "woof"
}

func Say(s Speaker) string {
    if s == nil {
        return ""
    }
    return s.Speak()
}
```

`Dog` 不用写 `implements Speaker`。有 `Speak() string` 就算实现了。

## 命令

```bash
go run lessons/09-interface/demo.go
go test ./lessons/09-interface
```

## 练习

`Speaker` 有 `Speak() string`。`Say` 调用它并返回结果。`Dog.Speak` 返回 `woof`。
