# 7. 指针

`&x` 取地址，`*p` 取出指针指向的值。函数参数是值拷贝；要让函数改掉调用方的变量，就传指针。

不要把指针理解成 JS 的对象引用。结构体很大或需要修改原值时才用指针，不是默认到处加 `*`。

## 命令

```bash
go run lessons/07-pointer/demo.go
go test ./lessons/07-pointer
```

## 练习

`Swap` 交换两个 int 指针指向的值。
