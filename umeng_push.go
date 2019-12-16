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

type Result struct {
	Ret  string            `json:"ret"`
	Data map[string]string `json:"data"`
}

func (r *Result) IsSuccess() bool {
	if r.Ret == RetSuccess {
		return true
	}
	return false
}

func (r *Result) ErrorCode() string {
	if code, ok := r.Data["error_code"]; ok {
		return code
	}
	return ""
}

func (r *Result) ErrorMsg() string {
	if msg, ok := r.Data["error_msg"]; ok {
		return msg
	}
	return ""
}

func Post(url string, data []byte) (result Result, err error) {
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

	body, _ := ioutil.ReadAll(rsp.Body)
	err = json.Unmarshal(body, &result)

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

func (u *UmengPush) Send(param *Param) (result Result, err error) {
	param.AppKey = u.AppKey
	param.Timestamp = strconv.Itoa(int(time.Now().Unix()))
	data, err := json.Marshal(param)
	if err != nil {
		return
	}

	url := UrlSign(MessageSend, string(data), u.AppMasterKey)
	result, err = Post(url, data)
	if err != nil {
		return
	} else if !result.IsSuccess() {
		err = errors.New(fmt.Sprintf("error_code=%s;error_msg=%s", result.ErrorCode(), result.ErrorMsg()))
	}
	return
}
