# 18. JWT

登录成功后签发 token，后面的请求在 `Authorization: Bearer <token>` 里带上。服务端用同一个密钥解析，并检查过期时间。

不要把密码放进 token。token 里通常只放用户 id。`Claims` 结构体已经在 `exercise.go` 里。

## 怎么写

签发：

```go
claims := Claims{
    UserID: userID,
    RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(exp),
    },
}
return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
```

解析：`jwt.ParseWithClaims(raw, &Claims{}, keyFunc)`。`keyFunc` 里确认算法是 HS256，再 `return []byte(secret), nil`。成功后把 `parsed.Claims` 转成 `*Claims`，返回 `UserID`。

## 命令

```bash
go run lessons/18-jwt/demo.go
go test ./lessons/18-jwt
```

## 练习

`Sign` 用 HS256 签发，claims 里放 `uid` 和过期时间。`UserID` 用同一密钥解析，过期或签名不对时返回错误。
