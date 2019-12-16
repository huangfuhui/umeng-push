package umeng_push

// Android调用参数说明
// {
//    "appkey":"xx",        // 必填，应用唯一标识
//    "timestamp":"xx",    // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
//    "type":"xx",        // 必填，消息发送类型,其值可以为:
//                        //   unicast-单播
//                        //   listcast-列播，要求不超过500个device_token
//                        //   filecast-文件播，多个device_token可通过文件形式批量发送
//                        //   broadcast-广播
//                        //   groupcast-组播，按照filter筛选用户群, 请参照filter参数
//                        //   customizedcast，通过alias进行推送，包括以下两种case:
//                        //     - alias: 对单个或者多个alias进行推送
//                        //     - file_id: 将alias存放到文件后，根据file_id来推送
//    "device_tokens":"xx",    // 当type=unicast时, 必填, 表示指定的单个设备
//                            // 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
//    "alias_type": "xx",    // 当type=customizedcast时, 必填
//                        // alias的类型, alias_type可由开发者自定义, 开发者在SDK中
//                        // 调用setAlias(alias, alias_type)时所设置的alias_type
//    "alias":"xx",        // 当type=customizedcast时, 选填(此参数和file_id二选一)
//                        // 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
//                        // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
//    "file_id":"xx",    // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
//                    // 当type=customizedcast时，选填(此参数和alias二选一)
//                    //   file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应
//                    //   的alias_type必须和接口参数alias_type一致。
//                    // 使用文件播需要先调用文件上传接口获取file_id，参照"文件上传"
//    "filter":{},    // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G。
//                    // filter的内容长度最大为3000B）
//    "payload": {    // 必填，JSON格式，具体消息内容(Android最大为1840B)
//        "display_type":"xx",    // 必填，消息类型: notification(通知)、message(消息)
//        "body": {    // 必填，消息体。
//                // 当display_type=message时，body的内容只需填写custom字段。
//                // 当display_type=notification时，body包含如下参数:
//            // 通知展现内容:
//            "ticker":"xx",    // 必填，通知栏提示文字
//            "title":"xx",    // 必填，通知标题
//            "text":"xx",    // 必填，通知文字描述
//
//            // 自定义通知图标:
//            "icon":"xx",    // 可选，状态栏图标ID，R.drawable.[smallIcon]，
//            // 如果没有，默认使用应用图标。
//            // 图片要求为24*24dp的图标，或24*24px放在drawable-mdpi下。
//            // 注意四周各留1个dp的空白像素
//            "largeIcon":"xx",    // 可选，通知栏拉开后左侧图标ID，R.drawable.[largeIcon]，
//            // 图片要求为64*64dp的图标，
//            // 可设计一张64*64px放在drawable-mdpi下，
//            // 注意图片四周留空，不至于显示太拥挤
//            "img": "xx",    // 可选，通知栏大图标的URL链接。该字段的优先级大于largeIcon。
//                            // 该字段要求以http或者https开头。
//
//            // 自定义通知声音:
//            "sound": "xx",    // 可选，通知声音，R.raw.[sound]。
//                            // 如果该字段为空，采用SDK默认的声音，即res/raw/下的
//                            // umeng_push_notification_default_sound声音文件。如果
//                            // SDK默认声音文件不存在，则使用系统默认Notification提示音。
//
//            // 自定义通知样式:
//            "builder_id": xx,    // 可选，默认为0，用于标识该通知采用的样式。使用该参数时，
//                                // 开发者必须在SDK里面实现自定义通知栏样式。
//
//            // 通知到达设备后的提醒方式，注意，"true/false"为字符串
//            "play_vibrate":"true/false",    // 可选，收到通知是否震动，默认为"true"
//            "play_lights":"true/false",        // 可选，收到通知是否闪灯，默认为"true"
//            "play_sound":"true/false",        // 可选，收到通知是否发出声音，默认为"true"
//
//            // 点击"通知"的后续行为，默认为打开app。
//            "after_open": "xx",    // 可选，默认为"go_app"，值可以为:
//                                //   "go_app": 打开应用
//                                //   "go_url": 跳转到URL
//                                //   "go_activity": 打开特定的activity
//                                //   "go_custom": 用户自定义内容。
//            "url": "xx",    // 当after_open=go_url时，必填。
//                            // 通知栏点击后跳转的URL，要求以http或者https开头
//            "activity":"xx",    // 当after_open=go_activity时，必填。
//                                // 通知栏点击后打开的Activity
//            "custom":"xx"/{}    // 当display_type=message时, 必填
//                                // 当display_type=notification且
//                                // after_open=go_custom时，必填
//                                // 用户自定义内容，可以为字符串或者JSON格式。
//        },
//        extra:{    // 可选，JSON格式，用户自定义key-value。只对"通知"
//                // (display_type=notification)生效。
//                // 可以配合通知到达后，打开App/URL/Activity使用。
//            "key1": "value1",
//            "key2": "value2",
//            ...
//        }
//    },
//    "policy":{    // 可选，发送策略
//        "start_time":"xx",    // 可选，定时发送时，若不填写表示立即发送。
//                            // 定时发送时间不能小于当前时间
//                            // 格式: "yyyy-MM-dd HH:mm:ss"。
//                            // 注意，start_time只对任务类消息生效。
//        "expire_time":"xx",    // 可选，消息过期时间，其值不可小于发送时间或者
//                            // start_time(如果填写了的话)，
//                            // 如果不填写此参数，默认为3天后过期。格式同start_time
//        "max_send_num": xx,    // 可选，发送限速，每秒发送的最大条数。最小值1000
//                            // 开发者发送的消息如果有请求自己服务器的资源，可以考虑此参数。
//        "out_biz_no": "xx"    // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送。
//                            // 有些情况下（例如网络异常）开发者可能会重复调用API导致
//                            // 消息多次下发到客户端。如果需要处理这种情况，可以考虑此参数。
//                            // 注意, out_biz_no只对任务类消息生效。
//    },
//    "production_mode":"true/false",    // 可选，正式/测试模式。默认为true
//                                    // 测试模式只对“广播”、“组播”类消息生效，其他类型的消息任务（如“文件播”）不会走测试模式
//                                    // 测试模式只会将消息发给测试设备。测试设备需要到web上添加。
//                                    // Android: 测试设备属于正式设备的一个子集。
//    "description": "xx",    // 可选，发送消息描述，建议填写。
//    //系统弹窗，只有display_type=notification生效
//    "mipush": "true/false",    // 可选，默认为false。当为true时，表示MIUI、EMUI、Flyme系统设备离线转为系统下发
//    "mi_activity": "xx",    // 可选，mipush值为true时生效，表示走系统通道时打开指定页面acitivity的完整包路径。
//}
//
// *********************************************************************************************************************
//
// IOS调用参数说明
// {
//  "appkey":"xx",    // 必填，应用唯一标识
//  "timestamp":"xx", // 必填，时间戳，10位或者13位均可，时间戳有效期为10分钟
//  "type":"xx",      // 必填，消息发送类型,其值可以为:
//                    //   unicast-单播
//                    //   listcast-列播，要求不超过500个device_token
//                    //   filecast-文件播，多个device_token可通过文件形式批量发送
//                    //   broadcast-广播
//                    //   groupcast-组播，按照filter筛选用户群, 请参照filter参数
//                    //   customizedcast，通过alias进行推送，包括以下两种case:
//                    //     - alias: 对单个或者多个alias进行推送
//                    //     - file_id: 将alias存放到文件后，根据file_id来推送
//  "device_tokens":"xx", // 当type=unicast时, 必填, 表示指定的单个设备
//                        // 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
//  "alias_type": "xx", // 当type=customizedcast时, 必填
//                      // alias的类型, alias_type可由开发者自定义, 开发者在SDK中
//                      // 调用setAlias(alias, alias_type)时所设置的alias_type
//  "alias":"xx", // 当type=customizedcast时, 选填(此参数和file_id二选一)
//                // 开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
//                // 在SDK中调用setAlias(alias, alias_type)时所设置的alias
//  "file_id":"xx", // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
//                  // 当type=customizedcast时，选填(此参数和alias二选一)
//                  //   file内容为多条alias，以回车符分隔。注意同一个文件内的alias所对应
//                  //   的alias_type必须和接口参数alias_type一致。
//                  // 使用文件播需要先调用文件上传接口获取file_id，参照"2.4文件上传接口"
//  "filter":{}, // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G。
//  "payload":   // 必填，JSON格式，具体消息内容(iOS最大为2012B)
//  {
//    "aps":      // 必填，严格按照APNs定义来填写
//    {
//        "alert":""/{ // 当content-available=1时(静默推送)，可选; 否则必填。
//                     // 可为JSON类型和字符串类型
//            "title":"title",
//            "subtitle":"subtitle",
//            "body":"body"
//        }
//        "badge": xx,           // 可选
//        "sound": "xx",         // 可选
//        "content-available":1  // 可选，代表静默推送
//        "category": "xx",      // 可选，注意: ios8才支持该字段。
//    },
//    "key1":"value1",       // 可选，用户自定义内容, "d","p"为友盟保留字段，
//                           // key不可以是"d","p"
//    "key2":"value2",
//    ...
//  },
//  "policy":               // 可选，发送策略
//  {
//      "start_time":"xx",   // 可选，定时发送时间，若不填写表示立即发送。
//                           // 定时发送时间不能小于当前时间
//                           // 格式: "yyyy-MM-dd HH:mm:ss"。
//                           // 注意，start_time只对任务生效。
//      "expire_time":"xx",  // 可选，消息过期时间，其值不可小于发送时间或者
//                           // start_time(如果填写了的话),
//                           // 如果不填写此参数，默认为3天后过期。格式同start_time
//      "out_biz_no": "xx"   // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送。
//                           // 有些情况下（例如网络异常）开发者可能会重复调用API导致
//                           // 消息多次下发到客户端。如果需要处理这种情况，可以考虑此参数。
//                           // 注意，out_biz_no只对任务生效。
//      "apns_collapse_id": "xx" // 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展示
//                               // 最新的一条，字段长度不得超过64bytes
//  },
//  "production_mode":"true/false" // 可选，正式/测试模式。默认为true
//                                 // 测试模式只对“广播”、“组播”类消息生效，其他类型的消息任务（如“文件播”）不会走测试模式
//                                 // 测试模式只会将消息发给测试设备。测试设备需要到web上添加。
//  "description": "xx"      // 可选，发送消息描述，建议填写。
//}
type Param struct {
	AppKey       string      `json:"appkey"`
	Timestamp    string      `json:"timestamp"`
	Types        string      `json:"type"`
	DeviceTokens string      `json:"device_tokens,omitempty"`
	AliasType    string      `json:"alias_type,omitempty"`
	Alias        string      `json:"alias,omitempty"`
	FileId       string      `json:"file_id,omitempty"`
	Filter       string      `json:"filter,omitempty"`
	Payload      interface{} `json:"payload"`
	Policy       struct {
		StartTime  string `json:"start_time,omitempty"`
		ExpireTime string `json:"expire_time,omitempty"`
		MaxSendNum int64  `json:"max_send_num,omitempty"`
		OutBizNo   string `json:"out_biz_no,omitempty"`
	} `json:"policy,omitempty"`
	ProductionMode string `json:"production_mode,omitempty"`
	Description    string `json:"description,omitempty"`
	Mipush         string `json:"mipush,omitempty"`
	MiActivity     string `json:"mi_activity,omitempty"`
}

type AndroidPayload struct {
	DisplayType string `json:"display_type"`
	Body        struct {
		Ticker      string            `json:"ticker"`
		Title       string            `json:"title"`
		Text        string            `json:"text"`
		Icon        string            `json:"icon,omitempty"`
		LargeIcon   string            `json:"largeIcon,omitempty"`
		Img         string            `json:"img,omitempty"`
		Sound       string            `json:"sound,omitempty"`
		BuilderId   string            `json:"builder_id,omitempty"`
		PlayVibrate string            `json:"play_vibrate,omitempty"`
		PlayLights  string            `json:"play_lights,omitempty"`
		PlaySound   string            `json:"play_sound,omitempty"`
		AfterOpen   string            `json:"after_open,omitempty"`
		Url         string            `json:"url,omitempty"`
		Activity    string            `json:"activity,omitempty"`
		Custom      string            `json:"custom,omitempty"`
		Extra       map[string]string `json:"extra,omitempty"`
	} `json:"body"`
}

type IosPayload struct {
	Aps struct {
		Alert struct {
			Title    string `json:"title,omitempty"`
			Subtitle string `json:"subtitle,omitempty"`
			Body     string `json:"body,omitempty"`
		} `json:"alert,omitempty"`
		Badge            int64  `json:"badge,omitempty"`
		Sound            string `json:"sound,omitempty"`
		ContentAvailable int64  `json:"content-available,omitempty"`
		Category         string `json:"category,omitempty"`
	} `json:"aps"`
}
