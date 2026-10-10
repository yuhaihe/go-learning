# 4. 函数与多返回值

函数可以返回多个值。习惯是最后一个是 `error`，调用方立刻判断。

不要用异常当正常分支。Go 没有 try/catch。

## 命令

```bash
go run lessons/04-functions/demo.go
go test ./lessons/04-functions
```

## 练习

`Divide(a, b)` 做整数除法。`b == 0` 时返回错误，商返回 `0`。
