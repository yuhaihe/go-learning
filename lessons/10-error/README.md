# 10. 错误

正常失败返回 `error`，不要 `panic`。`panic` 只留给程序无法继续的情况，例如启动时配置缺失。

用 `fmt.Errorf("...: %w", err)` 包一层，调用方用 `errors.Is` 判断是不是某一种错误。

## 命令

```bash
go run lessons/10-error/demo.go
go test ./lessons/10-error
```

## 练习

`ParsePositive` 把字符串解析成正整数。不是数字时包住 `strconv` 的错误。小于等于 0 时包住 `ErrNotPositive`。
