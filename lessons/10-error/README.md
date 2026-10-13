# 10. 错误

正常失败返回 `error`，不要 `panic`。`panic` 只留给程序无法继续的情况，例如启动时配置缺失。

用 `fmt.Errorf("...: %w", err)` 包一层，调用方用 `errors.Is` 判断是不是某一种错误。`%w` 才能被 `errors.Is` 认出来；写成 `%v` 就包不住。

## 这一课的词

`errors.New("...")` 做一个固定文案的错误。要能被认出来，必须始终用同一个变量，比如练习里的 `ErrNotPositive`。重新 `errors.New` 一次，文案一样也对不上。

`fmt.Errorf` 里的 `%w` 把旧错误包进去，关系还在。`%v` 只把文字拼进新错误，关系断了。`%q` 会给字符串加上引号，`"x"` 印出来是 `"x"`。

`errors.Is(err, ErrNotPositive)` 顺着包进去的链，看是不是这个错误。`errors.As` 看链里有没有某一种类型。测试用它认出 `strconv` 的数字解析错误。

## 怎么写

```go
import "strconv"

n, err := strconv.Atoi(s) // "12" -> 12；"x" -> err
if err != nil {
    return 0, fmt.Errorf("parse %q: %w", s, err)
}
if n <= 0 {
    return 0, fmt.Errorf("parse %q: %w", s, ErrNotPositive)
}
```

`ErrNotPositive` 已经在 `exercise.go` 里声明好了。

## 命令

```bash
go run lessons/10-error/demo.go
go test ./lessons/10-error
```

## 练习

`ParsePositive` 把字符串解析成正整数。不是数字时包住 `strconv` 的错误。小于等于 0 时包住 `ErrNotPositive`。
