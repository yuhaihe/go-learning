# 5. 数组与 slice

数组长度是类型的一部分：`[3]int` 和 `[4]int` 不是同一种类型。

slice 是数组的视图，长度可变。`append` 可能换一块新底层数组，所以要把返回值接回去：`s = append(s, x)`。

JS 数组更接近 slice，不是 Go 的数组。

## 命令

```bash
go run lessons/05-slice/demo.go
go test ./lessons/05-slice
```

## 练习

`Unique` 去掉重复的 int，保留第一次出现的顺序。
