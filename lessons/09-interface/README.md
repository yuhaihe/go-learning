# 9. 接口

接口只描述方法。类型不用声明「我实现了某接口」，方法对上就算实现。

接口值由「具体类型」和「具体值」组成。业务代码依赖接口，测试时换一个假的实现。

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
