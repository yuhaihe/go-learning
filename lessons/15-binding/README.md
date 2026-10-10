# 15. 参数绑定与校验

`ShouldBindJSON` 把请求体填进结构体。字段上的 `binding` 标签负责校验，例如 `required`、`min`、`max`。

校验失败要返回 400，不要继续执行业务。

## 命令

```bash
go test ./lessons/15-binding
```

## 练习

`CreateNote` 的 `title` 必填且不超过 100 个字符，`content` 不超过 2000。`ParseCreateNote` 绑定 JSON，失败时返回错误。
