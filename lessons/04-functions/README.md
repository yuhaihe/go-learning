# 4. 函数与多返回值

函数可以返回多个值。习惯是最后一个是 `error`，调用方立刻判断。

不要用异常当正常分支。Go 没有 try/catch。

## 怎么写

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}
```

调用时先看错误。不要的返回值用 `_`：

```go
q, err := divide(10, 2)
if err != nil {
    // 处理错误，不要用 q
}
_, err = divide(1, 0) // 商不需要时用 _
```

`:=` 是第一次声明。后面再赋值用 `=`，变量必须已经存在。

## 命令

```bash
go run lessons/04-functions/demo.go
go test ./lessons/04-functions
```

## 练习

`Divide(a, b)` 做整数除法。`b == 0` 时返回错误，商返回 `0`。
