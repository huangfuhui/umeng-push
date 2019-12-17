package umeng_push

// 消息发送-调用返回
type SendResult struct {
	Ret  string `json:"ret"` // SUCCESS/FAIL
	Data struct {
		// 当"ret"为"SUCCESS"时,包含如下参数:
		MsgId  string `json:"msg_id"`  // 单播类消息(type为unicast、listcast、customizedcast且不带file_id)返回
		TaskId string `json:"task_id"` // 任务类消息(type为broadcast、groupcast、filecast、customizedcast且file_id不为空)返回

		// 当"ret"为"FAIL"时,包含如下参数:
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

// 状态查询-调用返回
type StatusResult struct {
	Ret  string `json:"ret"` // SUCCESS/FAIL
	Data struct {
		// 当"ret"为"SUCCESS"时,包含如下参数:
		TaskId       string `json:"task_id"`       // 任务类消息(type为broadcast、groupcast、filecast、customizedcast且file_id不为空)返回
		Status       int64  `json:"status"`        // 消息状态: 0-排队中, 1-发送中,2-发送完成,3-发送失败,4-消息被撤销,5-消息过期, 6-筛选结果为空,7-定时任务尚未开始处理
		SentCount    int64  `json:"sent_count"`    // 消息收到数
		OpenCount    int64  `json:"open_count"`    // 打开数
		DismissCount int64  `json:"dismiss_count"` // 忽略数
		TotalCount   int64  `json:"total_count"`   // 投递APNs设备数

		// 当"ret"为"FAIL"时,包含如下参数:
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (r *StatusResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}

// 消息取消-调用返回
type CancelResult struct {
	Ret  string `json:"ret"` // SUCCESS/FAIL
	Data struct {
		// 当"ret"为"SUCCESS"时,包含如下参数:
		TaskId string `json:"task_id"` // 任务类消息(type为broadcast、groupcast、filecast、customizedcast且file_id不为空)返回

		// 当"ret"为"FAIL"时,包含如下参数:
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

// 文件上传-调用返回
type UploadResult struct {
	Ret  string `json:"ret"` // SUCCESS/FAI
	Data struct {
		// 当"ret"为"SUCCESS"时,包含如下参数:
		FileId string `json:"file_id"`

		// 当"ret"为"FAIL"时,包含如下参数:
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

// 管理用户自定义标签-调用返回
type TagResult struct {
	Ret  string `json:"ret"` // SUCCESS/FAIL
	Data struct {
		// 当"ret"为"FAIL"时,包含如下参数:
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	} `json:"data"`
}

func (r *TagResult) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}
