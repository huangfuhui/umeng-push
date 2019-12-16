package umeng_push

//{
//    "ret":"SUCCESS/FAIL",
//    "data": {
//        // 当"ret"为"SUCCESS"时，包含如下参数:
//        // 单播类消息(type为unicast、listcast、customizedcast且不带file_id)返回：
//        "msg_id":"xx"
//
//        // 任务类消息(type为broadcast、groupcast、filecast、customizedcast且file_id不为空)返回：
//        "task_id":"xx"
//
//        // 当"ret"为"FAIL"时,包含如下参数:
//        "error_code":"xx",    // 错误码，详见附录I
//        "error_msg":"xx"    // 错误信息
//    }
//}
type SendResult struct {
	Ret  string `json:"ret"`
	Data struct {
		MsgId     string `json:"msg_id"`
		TaskId    string `json:"task_id"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (r *SendResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}

//{
//    "ret":"SUCCESS/FAIL",
//    "data": {
//        // 当"ret"为"SUCCESS"时，包含如下参数:
//        "task_id":"xx",
//        "status": xx,    // 消息状态: 0-排队中, 1-发送中，2-发送完成，3-发送失败，4-消息被撤销，
//                        // 5-消息过期, 6-筛选结果为空，7-定时任务尚未开始处理
//        // Android消息，包含以下参数
//        "sent_count":xx,    // 消息收到数
//        "open_count":xx,    // 打开数
//        "dismiss_count":xx    // 忽略数
//
//        // iOS消息，包含以下参数
//        "total_count": xx,    // 投递APNs设备数
//        "sent_count": xx,    // APNs返回SUCCESS的设备数
//        "open_count": xx    // 打开数
//
//        // 当"ret"为"FAIL"时，包含参数如下:
//        "error_code": "xx",    // 错误码详见附录I。
//        "error_msg": "xx"    // 错误码详见附录I。
//      }
//}
type StatusResult struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskId       string `json:"task_id"`
		Status       int64  `json:"status"`
		SentCount    int64  `json:"sent_count"`
		OpenCount    int64  `json:"open_count"`
		DismissCount int64  `json:"dismiss_count"`
		TotalCount   int64  `json:"total_count"`
		ErrorCode    string `json:"error_code"`
		ErrorMsg     string `json:"error_msg"`
	} `json:"data"`
}

func (r *StatusResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}

//{
//    "ret":"SUCCESS/FAIL",
//    "data": {
//        // 当"ret"为"SUCCESS"时
//        "task_id":"xx"
//
//        // 当"ret"为"FAIL"时，包含参数如下:
//        "error_code": "xx",    // 错误码
//        "error_msg": "xx"    // 错误详情
//    }
//}
type CancelResult struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskId    string `json:"task_id"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (r *CancelResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}

//{
//    "ret":"SUCCESS/FAIL",
//    "data": {
//        // 当"ret"为"SUCCESS"时
//        "file_id":"xx"
//
//        // 当"ret"为"FAIL"时，包含参数如下:
//        "error_code": "xx",    //错误码
//        "error_msg": "xx"    // 错误详情
//    }
//}
type UploadResult struct {
	Ret  string `json:"ret"`
	Data struct {
		FileId    string `json:"file_id"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (r *UploadResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}
