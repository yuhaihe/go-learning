# 15. 参数绑定与校验

`ShouldBindJSON` 把请求体填进结构体。字段上的 `binding` 标签负责校验，例如 `required`、`min`、`max`。

校验失败要返回 400，不要继续执行业务。结构体和 `binding` 标签已经写在 `exercise.go` 里。

## 这一课的词

字段后面反引号里的内容叫结构体标签，给库看，不是字段的值。

`json:"title"` 表示 JSON 里的键名是 `title`。`binding:"required,max=100"` 表示必填，并且最长 100 个字符。

`ShouldBindJSON(&req)` 读取请求体，按标签填进 `req` 并校验。失败时 `err != nil`，把这个错误返回去。测试会把它当成 400。

## 怎么写

```go
var req CreateNote
if err := c.ShouldBindJSON(&req); err != nil {
    return CreateNote{}, err
}
return req, nil
```

## 命令

```bash
go test ./lessons/15-binding
```

## 练习

`CreateNote` 的 `title` 必填且不超过 100 个字符，`content` 不超过 2000。`ParseCreateNote` 绑定 JSON，失败时返回错误。
