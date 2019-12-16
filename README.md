# umeng-push
<a href="https://developer.umeng.com/docs/66632/detail/68343">友盟推送</a>友盟推送服务端SDK，默认全部接口使用HTTPS协议。

##### 一、使用示例
```
// 初始化客户端
var appKey = "your app_key"
var appMasterKey = "your app_master_key"
umengPush = NewUmengPush(appKey, appMasterKey)

// 根据业务装填参数
param := &SendParam{}

// 请求调用
result, err := umengPush.Send(param)
if err != nil {
	log.Fatal(err)
}
log.Println(result.Data.TaskId)
```

##### 二、已实现接口列表
- 消息发送
- 任务类消息状态查询
- 任务类消息取消
- 文件上传
