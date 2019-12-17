// 友盟推送友盟推送服务端SDK，默认全部接口使用HTTPS协议。
//
// 下载安装
//	$ go get -u github.com/huangfuhui/umeng-push
//
// 使用示例
//	// 初始化客户端
//	var appKey = "your app_key"
//	var appMasterKey = "your app_master_key"
//	umengPush := umeng_push.NewUmengPush(appKey, appMasterKey)
//
//	// 根据业务装填参数
//	param := &umeng_push.SendParam{}
//
//	// 请求调用
//	result, err := umengPush.Send(param)
//		if err != nil {
//		log.Fatal(err)
//	}
package umeng_push
