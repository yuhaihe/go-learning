# 2. 变量、常量、零值

`var n int` 不用赋值也有值，int 是 `0`，string 是 `""`，bool 是 `false`。这叫零值。

`:=` 是短声明，只能写在函数内部。常量用 `const`，编译期就定死。

JS 里未赋值往往是 `undefined`。Go 的变量总有一个可用的零值。

## 怎么写

`:= 0` 仍然是赋值。这课要的是只声明、不写初值：

```go
var n int    // 0
var s string // ""
var ok bool  // false
return n, s, ok
```

## 命令

```bash
go run lessons/02-variables/demo.go
go test ./lessons/02-variables
```

## 练习

`ZeroValues` 只声明局部变量，不要赋值，把 int、string、bool 的零值返回出去。
