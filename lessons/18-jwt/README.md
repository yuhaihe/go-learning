# 18. JWT

登录成功后签发一段 token，后面的请求在 `Authorization: Bearer <token>` 里带上。服务端用同一个密钥解析，并检查过期时间。

不要把密码放进 token。token 里通常只放用户 id。`Claims` 结构体已经在 `exercise.go` 里。

## 这一课的词

token 是一串签过名的文本。客户端没法偷偷改里面的用户 id，改了签名就对不上。

HS256 是签名算法，签发和解析用同一把密钥 `secret`。`Claims` 是 token 里的内容，这里有用户 id 和过期时间 `exp`。

`SignedString` 签出最终那串 token。`ParseWithClaims` 用密钥把它解回来。过期或签名不对都会返回错误。

`Bearer` 只是请求头里的一个词，后面跟 token 本身。这课的函数只负责签发和解析，不处理请求头。

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
