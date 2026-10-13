# 5. 数组与 slice

## 数组 vs slice

| | 数组 `[3]int` | slice `[]int` |
|---|---------------|---------------|
| 长度 | 固定，写进类型 | 可变 |
| 字面量 | `[3]int{1,2,3}` | `[]int{1,2,3}` 或 `make([]int, 0)` |
| 追加 | 不能 append | `s = append(s, x)` |

JS 里的「数组」更接近 Go 的 **slice**。

## 怎么遍历 slice

**方式 1：`for range`（最常用）**

```go
nums := []int{1, 2, 3}
for i, v := range nums {
    // i 下标，v 当前元素
    _ = i
    _ = v
}
// 只要值不要下标：
for _, v := range nums {
    _ = v
}
```

**方式 2：按下标**

```go
for i := 0; i < len(nums); i++ {
    n := nums[i]
    _ = n
}
```

不要用变量名 `len` 去接 `len(nums)`，会遮住内置函数 `len`。

## slice 常用操作（不是 `.方法()`，是函数）

Go 的 slice **没有** JS 那种 `push` / `filter`。记这几个：

```go
n := len(s)           // 元素个数
c := cap(s)           // 容量（demo 会打印）
s = append(s, 10)     // 末尾加一个，必须 s = append(...)
s = append(s, 20, 30) // 可加多个
```

去重、过滤没有内置一行搞定，一般是：**遍历 + 判断 + append 到新 slice**。

## 做 `Unique` 的思路（不贴完整代码）

1. 准备一个**结果 slice**（可先 `make([]int, 0)` 或 `var out []int`）。
2. 准备一个 **map**，记录「某个数是否已经出现过」。`map[int]bool` 就够：`true` 表示见过。  
   `map[int]struct{}` 是同一种记法。`struct{}` 是空结构体，不占数据，只用来表示「这个键在」。`struct{}{}` 是它的一个空值。用 `bool` 更直观。
3. `for _, n := range nums`：若没见过，append 到结果，并在 map 里标记见过。
4. 返回结果 slice。输入 `nil` 时应返回空 slice（长度 0），不是 `nil` 也可以，测试要求空 slice。

## 命令

```bash
go run lessons/05-slice/demo.go
go test ./lessons/05-slice
```

`exercise.go` 的包名是 **`slicex`**，不要写 `func main()`。调试请改 `demo.go` 或跑测试。

## 练习

实现 `Unique`：去掉重复的 int，保留第一次出现的顺序。
