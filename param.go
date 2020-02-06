package umeng_push

// 消息发送调用参数
type SendParam struct {
	AppKey    string `json:"appkey"`    // 必填,应用唯一标识
	Timestamp string `json:"timestamp"` // 必填,时间戳,10位或者13位均可,时间戳有效期为10分钟

	// 必填,消息发送类型,其值可以为:
	//   unicast-单播
	//   listcast-列播,要求不超过500个device_token
	//   filecast-文件播,多个device_token可通过文件形式批量发送
	//   broadcast-广播
	//   groupcast-组播,按照filter筛选用户群, 请参照filter参数
	//   customizedcast,通过alias进行推送,包括以下两种case:
	//     - alias: 对单个或者多个alias进行推送
	//     - file_id: 将alias存放到文件后,根据file_id来推送
	Types string `json:"type"`

	// 当type=unicast时, 必填, 表示指定的单个设备
	// 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
	DeviceTokens string `json:"device_tokens,omitempty"`

	// 当type=customizedcast时, 必填
	// alias的类型, alias_type可由开发者自定义, 开发者在SDK中
	// 调用setAlias(alias, alias_type)时所设置的alias_type
	AliasType string `json:"alias_type,omitempty"`

	// 当type=customizedcast时, 选填(此参数和file_id二选一)
	// 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
	// 在SDK中调用setAlias(alias, alias_type)时所设置的alias
	Alias string `json:"alias,omitempty"`

	// 当type=filecast时,必填,file内容为多条device_token,以回车符分割
	// 当type=customizedcast时,选填(此参数和alias二选一)
	//   file内容为多条alias,以回车符分隔,注意同一个文件内的alias所对应
	//   的alias_type必须和接口参数alias_type一致
	// 使用文件播需要先调用文件上传接口获取file_id,参照"2.4文件上传接口"
	FileId string `json:"file_id,omitempty"`

	Filter  string      `json:"filter,omitempty"` // 当type=groupcast时,必填,用户筛选条件,如用户标签、渠道等,参考附录G
	Payload interface{} `json:"payload"`          // 必填,JSON格式,具体消息内容(iOS最大为2012B,Android最大为1840B)
	Policy  struct {
		// 可选,定时发送时间,若不填写表示立即发送
		// 定时发送时间不能小于当前时间
		// 格式: "yyyy-MM-dd HH:mm:ss"
		// 注意,start_time只对任务生效
		StartTime string `json:"start_time,omitempty"`

		// 可选,消息过期时间,其值不可小于发送时间或者
		// start_time(如果填写了的话),
		// 如果不填写此参数,默认为3天后过期,格式同start_time
		ExpireTime string `json:"expire_time,omitempty"`

		// 可选,发送限速,每秒发送的最大条数,最小值1000
		// 开发者发送的消息如果有请求自己服务器的资源,可以考虑此参数
		MaxSendNum int64 `json:"max_send_num,omitempty"`

		// 可选,开发者对消息的唯一标识,服务器会根据这个标识避免重复发送
		// 有些情况下（例如网络异常）开发者可能会重复调用API导致
		// 消息多次下发到客户端,如果需要处理这种情况,可以考虑此参数
		// 注意,out_biz_no只对任务生效
		OutBizNo string `json:"out_biz_no,omitempty"`

		// 可选,多条带有相同apns_collapse_id的消息,iOS设备仅展示
		// 最新的一条,字段长度不得超过64bytes
		ApnsCollapseId string `json:"apns_collapse_id,omitempty"`
	} `json:"policy,omitempty"` // 可选,发送策略

	// 可选,正式/测试模式,默认为true
	// 测试模式只对“广播”、“组播”类消息生效,其他类型的消息任务（如“文件播”）不会走测试模式
	// 测试模式只会将消息发给测试设备,测试设备需要到web上添加
	ProductionMode string `json:"production_mode,omitempty"`
	Description    string `json:"description,omitempty"` // 可选,发送消息描述,建议填写

	Mipush     string `json:"mipush,omitempty"`      // 可选,默认为false,当为true时,表示MIUI、EMUI、Flyme系统设备离线转为系统下发
	MiActivity string `json:"mi_activity,omitempty"` // 可选,mipush值为true时生效,表示走系统通道时打开指定页面acitivity的完整包路径

	ReceiptUrl  string `json:"receipt_url,omitempty"`  // 开发者接受数据的地址,最大长度256字节
	ReceiptType string `json:"receipt_type,omitempty"` // 回执数据类型,1:送达回执;2:点击回执;3:送达和点击回执,默认为3
}

type AndroidPayload struct {
	DisplayType string `json:"display_type"` // 必填,消息类型: notification(通知)、message(消息)

	// 必填,消息体
	// 当display_type=message时,body的内容只需填写custom字段
	// 当display_type=notification时,body包含如下参数:
	Body struct {
		Ticker string `json:"ticker"` // 必填,通知栏提示文字
		Title  string `json:"title"`  // 必填,通知标题
		Text   string `json:"text"`   // 必填,通知文字描述

		// 可选,状态栏图标ID,R.drawable.[smallIcon],
		// 如果没有,默认使用应用图标
		// 图片要求为24*24dp的图标,或24*24px放在drawable-mdpi下
		// 注意四周各留1个dp的空白像素
		Icon string `json:"icon,omitempty"`

		// 可选,通知栏拉开后左侧图标ID,R.drawable.[largeIcon],
		// 图片要求为64*64dp的图标,
		// 可设计一张64*64px放在drawable-mdpi下,
		// 注意图片四周留空,不至于显示太拥挤
		LargeIcon string `json:"largeIcon,omitempty"`

		// 可选,通知栏大图标的URL链接,该字段的优先级大于largeIcon
		// 该字段要求以http或者https开头
		Img string `json:"img,omitempty"`

		// 可选,通知声音,R.raw.[sound]
		// 如果该字段为空,采用SDK默认的声音,即res/raw/下的
		// umeng_push_notification_default_sound声音文件,如果
		// SDK默认声音文件不存在,则使用系统默认Notification提示音
		Sound string `json:"sound,omitempty"`

		BuilderId   string `json:"builder_id,omitempty"`   // 可选,默认为0,用于标识该通知采用的样式,使用该参数时,开发者必须在SDK里面实现自定义通知栏样式
		PlayVibrate string `json:"play_vibrate,omitempty"` // 可选,收到通知是否震动,默认为"true"
		PlayLights  string `json:"play_lights,omitempty"`  // 可选,收到通知是否闪灯,默认为"true"
		PlaySound   string `json:"play_sound,omitempty"`   // 可选,收到通知是否发出声音,默认为"true"

		// 点击"通知"的后续行为,默认为打开app
		// 可选,默认为"go_app",值可以为:
		//   "go_app": 打开应用
		//   "go_url": 跳转到URL
		//   "go_activity": 打开特定的activity
		//   "go_custom": 用户自定义内容
		AfterOpen string `json:"after_open,omitempty"`

		// 当after_open=go_url时,必填
		// 通知栏点击后跳转的URL,要求以http或者https开头
		Url string `json:"url,omitempty"`

		// 当after_open=go_activity时,必填
		// 通知栏点击后打开的Activity
		Activity string `json:"activity,omitempty"`

		// 当display_type=message时, 必填
		// 当display_type=notification且
		// after_open=go_custom时,必填
		Custom interface{} `json:"custom,omitempty"`
	} `json:"body"`

	// 可选,JSON格式,用户自定义key-value,只对"通知"
	// (display_type=notification)生效
	// 可以配合通知到达后,打开App/URL/Activity使用
	Extra interface{} `json:"extra,omitempty"`
}

type IosPayload struct {
	Aps struct {
		// 当content-available=1时(静默推送),可选; 否则必填
		// 可为JSON类型和字符串类型
		Alert struct {
			Title    string `json:"title,omitempty"`
			Subtitle string `json:"subtitle,omitempty"`
			Body     string `json:"body,omitempty"`
		} `json:"alert,omitempty"`
		Badge            int64  `json:"badge,omitempty"`
		Sound            string `json:"sound,omitempty"`
		ContentAvailable int64  `json:"content-available,omitempty"` // 可选,代表静默推送
		Category         string `json:"category,omitempty"`          // 可选,注意: ios8才支持该字段
	} `json:"aps"` // 必填,严格按照APNs定义来填写

	// "key1":"value1",       // 可选,用户自定义内容, "d","p"为友盟保留字段, key不可以是"d","p"
	// "key2":"value2",
	// ...
}

// 状态查询调用参数
type StatusParam struct {
	AppKey    string `json:"appkey"`    // 必填, 应用唯一标识
	Timestamp string `json:"timestamp"` // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	TaskId    string `json:"task_id"`   // 必填, 消息发送时, 从返回消息中获取的task_id
}

// 任务类消息取消调用参数
type CancelParam struct {
	AppKey    string `json:"appkey"`    // 必填, 应用唯一标识
	Timestamp string `json:"timestamp"` // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	TaskId    string `json:"task_id"`   // 必填, 消息发送时, 从返回消息中获取的task_id
}

// 文件上传调用参数
type UploadParam struct {
	AppKey    string `json:"appkey"`    // 必填, 应用唯一标识
	Timestamp string `json:"timestamp"` // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	Content   string `json:"content"`   // 必填, 文件内容, 多个device_token/alias请用回车符"\n"分隔
}

// 给设备打标签调用参数
type TagAddParam struct {
	AppKey       string `json:"appkey"`                  // 必填, 应用唯一标识
	Timestamp    string `json:"timestamp"`               // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	DeviceTokens string `json:"device_tokens,omitempty"` // 单个device_token
	Tag          string `json:"tag"`                     // 要添加的标签,如果有多个,以英文逗号分隔
}

// 查询设备标签列表调用参数
type TagListParam struct {
	AppKey       string `json:"appkey"`                  // 必填, 应用唯一标识
	Timestamp    string `json:"timestamp"`               // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	DeviceTokens string `json:"device_tokens,omitempty"` // 只支持一个device_token
}

// 设置设备标签调用参数
type TagSetParam struct {
	AppKey       string `json:"appkey"`                  // 必填, 应用唯一标识
	Timestamp    string `json:"timestamp"`               // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	DeviceTokens string `json:"device_tokens,omitempty"` // 单个device_token
	Tag          string `json:"tag"`                     // 要添加的标签,如果有多个,以英文逗号分隔
}

// 删除设备标签调用参数
type TagDeleteParam struct {
	AppKey       string `json:"appkey"`                  // 必填, 应用唯一标识
	Timestamp    string `json:"timestamp"`               // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	DeviceTokens string `json:"device_tokens,omitempty"` // 只支持一个device_token
	Tag          string `json:"tag"`
}

// 清除设备标签调用参数
type TagClearParam struct {
	AppKey       string `json:"appkey"`                  // 必填, 应用唯一标识
	Timestamp    string `json:"timestamp"`               // 必填, 时间戳,10位或者13位均可,时间戳有效期为10分钟
	DeviceTokens string `json:"device_tokens,omitempty"` // 只支持一个device_token
}
