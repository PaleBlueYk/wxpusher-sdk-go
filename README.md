# wxpusher-client
> 个人定制修改，由于个人常用wxpusher，所以修改包名

[zjiecode/wxpusher-client](https://github.com/zjiecode/wxpusher-client)的 Go 语言版本

## 安装

```sh
go get -u github.com/PaleBlueYk/wxpusher-sdk-go
```

```sh
go get github.com/PaleBlueYk/wxpusher-sdk-go@v1.0.6
```

引入

```go
import (
	"github.com/PaleBlueYk/wxpusher-sdk-go"
	"github.com/PaleBlueYk/wxpusher-sdk-go/wxpusher"
)
```

## 发送消息

```go
msg := wxpusher.NewMessage(appToken).SetContent("测试").AddUId(uId)
msgArr, err := wxpusher.SendMessage(msg)
fmt.Println(msgArr, err)
```

## 查询状态

```go
status, err := wxpusher.QueryMessageStatus(2384429)
fmt.Println(status, err)
```

## 创建参数二维码

```go
qrcode := wxpusher.Qrcode{AppToken: appToken, Extra: "XX渠道用户"}
qrcodeResp, err := wxpusher.CreateQrcode(&qrcode)
fmt.Println(qrcodeResp, err)
```

## 查询 App 的关注用户

```go
result, err := wxpusher.QueryWxUser(appToken, 1, 20)
fmt.Println(result, err)
```
