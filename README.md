# 字节跳动小程序golang sdk

# 快速开始&demo
```go
// 创建字节小程序实例
app := microapp.New(microapp.Config{
    AppId:     "APPID",
    AppSecret: "SECRET",
})
```

### 登录
- [code2Session](https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session)
```go
res, err := auth.Code2SessionV2(app, auth.ApiCode2SessionV2Req{
    AnonymousCode: "anonymous_code",
    Code:          "code",
})
if err != nil {
    return nil, err
}
```





---
部分内容借鉴自：[github.com/fastwego/microapp](https://github.com/21888/microapp)

