package umeng_push

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

type Config struct {
	AppKey       string `json:"app_key"`
	AppMasterKey string `json:"app_master_key"`
	DeviceTokens string `json:"device_tokens"`
}

var (
	config Config

	umengPush *UmengPush
)

func TestMain(m *testing.M) {
	data, _ := ioutil.ReadFile("./debug/config.json")
	_ = json.Unmarshal(data, &config)
	m.Run()
}

func TestNewUmengPush(t *testing.T) {
	payload := &IosPayload{
		Aps: struct {
			Alert struct {
				Title    string `json:"title,omitempty"`
				Subtitle string `json:"subtitle,omitempty"`
				Body     string `json:"body,omitempty"`
			} `json:"alert,omitempty"`
			Badge            int64  `json:"badge,omitempty"`
			Sound            string `json:"sound,omitempty"`
			ContentAvailable int64  `json:"content-available,omitempty"`
			Category         string `json:"category,omitempty"`
		}{Alert: struct {
			Title    string `json:"title,omitempty"`
			Subtitle string `json:"subtitle,omitempty"`
			Body     string `json:"body,omitempty"`
		}{
			Title:    "标题",
			Subtitle: "子标题",
			Body:     "内容",
		}},
	}
	param := &SendParam{
		Types:        TypeUnicast,
		DeviceTokens: config.DeviceTokens,
		AliasType:    "",
		Alias:        "",
		FileId:       "",
		Filter:       "",
		Payload:      payload,
		Policy: struct {
			StartTime      string `json:"start_time,omitempty"`
			ExpireTime     string `json:"expire_time,omitempty"`
			MaxSendNum     int64  `json:"max_send_num,omitempty"`
			OutBizNo       string `json:"out_biz_no,omitempty"`
			ApnsCollapseId string `json:"apns_collapse_id,omitempty"`
		}{},
		ProductionMode: "true",
		Description:    "测试推送",
		Mipush:         "",
		MiActivity:     "",
	}

	umengPush = NewUmengPush(config.AppKey, config.AppMasterKey)
	result, err := umengPush.Send(param)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.Data)
}

func TestUmengPush_TagList(t *testing.T) {
	result, err := umengPush.TagList(config.DeviceTokens)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.Data)
}
