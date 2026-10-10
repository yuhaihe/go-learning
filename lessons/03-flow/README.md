# 3. 流程控制

Go 只有 `for`，没有 `while`。`for {}` 是死循环。`if` 和 `switch` 的条件不加括号。

`switch` 默认不会落到下一个分支，不用写 `break`。

## 怎么写

计数循环：

```go
for i := 0; i <= n; i++ {
    if i%2 == 0 {
        // i 是偶数
    }
}
```

按条件分支（没有 switch 变量）：

```go
switch {
case n < 0:
    return "neg"
case n == 0:
    return "zero"
default:
    return "pos"
}
```

`neg` / `zero` / `pos` 只是这题约定的三个字符串，不是关键字。

## 命令

```bash
go run lessons/03-flow/demo.go
go test ./lessons/03-flow
```

## 练习

`SumEven` 把 `0` 到 `n`（含）的偶数加起来。`n < 0` 时返回 `0`。

`Sign`：负数返回 `neg`，0 返回 `zero`，正数返回 `pos`。
