package umeng_push

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	MessageSend   = "https://msgapi.umeng.com/api/send"
	MessageStatus = "https://msgapi.umeng.com/api/status"
	MessageCancel = "https://msgapi.umeng.com/api/cancel"
	FileUpload    = "https://msgapi.umeng.com/upload"

	RetSuccess = "SUCCESS"
	RetFail    = "FAIL"

	TypeUnicast        = "unicast"
	TypeListcast       = "listcast"
	TypeFilecast       = "filecast"
	TypeBroadcast      = "broadcast"
	TypeGroupcast      = "groupcast"
	TypeCustomizedcast = "customizedcast"
)

func Post(url string, data []byte) (response []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	response, _ = ioutil.ReadAll(rsp.Body)

	return
}

type UmengPush struct {
	AppKey       string
	AppMasterKey string
}

func NewUmengPush(appKey, AppMasterKey string) *UmengPush {
	return &UmengPush{
		AppKey:       appKey,
		AppMasterKey: AppMasterKey,
	}
}

// 消息发送
func (u *UmengPush) Send(param *SendParam) (result SendResult, err error) {
	param.AppKey = u.AppKey
	param.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	url := UrlSign(MessageSend, string(data), u.AppMasterKey)
	response, err := Post(url, data)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if !result.IsSuccess() {
		err = errors.New(fmt.Sprintf("error_code=%s;error_msg=%s", result.Data.ErrorCode, result.Data.ErrorMsg))
	}
	return
}

// 任务类消息状态查询
func (u *UmengPush) Status(taskId string) (result StatusResult, err error) {
	param := StatusParam{
		AppKey:    u.AppKey,
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
		TaskId:    taskId,
	}
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	url := UrlSign(MessageStatus, string(data), u.AppMasterKey)
	response, err := Post(url, data)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if !result.IsSuccess() {
		err = errors.New(fmt.Sprintf("error_code=%s;error_msg=%s", result.Data.ErrorCode, result.Data.ErrorMsg))
	}
	return
}

// 任务类消息取消
func (u *UmengPush) Cancel(taskId string) (result CancelResult, err error) {
	param := CancelParam{
		AppKey:    u.AppKey,
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
		TaskId:    taskId,
	}
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	url := UrlSign(MessageCancel, string(data), u.AppMasterKey)
	response, err := Post(url, data)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if !result.IsSuccess() {
		err = errors.New(fmt.Sprintf("error_code=%s;error_msg=%s", result.Data.ErrorCode, result.Data.ErrorMsg))
	}
	return
}

// 文件上传
func (u *UmengPush) Upload(content string) (result UploadResult, err error) {
	param := UploadParam{
		AppKey:    u.AppKey,
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
		Content:   content,
	}
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	url := UrlSign(FileUpload, string(data), u.AppMasterKey)
	response, err := Post(url, data)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if !result.IsSuccess() {
		err = errors.New(fmt.Sprintf("error_code=%s;error_msg=%s", result.Data.ErrorCode, result.Data.ErrorMsg))
	}
	return
}
