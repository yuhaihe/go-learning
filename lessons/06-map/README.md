# 6. map

`map[string]int` 是键值表。读取不存在的键会得到零值，所以要用 `v, ok := m[k]` 区分「没有这个键」和「值就是零」。

map 是引用类型。作为参数传进函数，函数里的修改外面看得到。

## 命令

```bash
go run lessons/06-map/demo.go
go test ./lessons/06-map
```

## 练习

`Count` 统计每个单词出现的次数。
