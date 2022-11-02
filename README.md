# 字节跳动小程序golang sdk

# 快速开始&demo
```go
// 创建字节小程序实例
app := microapp.New(microapp.Config{
    AppId:     viper.GetString("APPID"),
    AppSecret: viper.GetString("SECRET"),
})
```

### 登录
- code2Session
```go
// 官方文档地址 https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session
// demo
params := url.Values{}// import "net/url"
params.Add("code", req.Code) // 这里虽然是用url拼接,但是post请求实现
session, err := auth.Code2SessionV2(app, params)
if err != nil {
    return nil, err
}
```





---
借鉴自：[github.com/fastwego/microapp](https://github.com/21888/microapp)

