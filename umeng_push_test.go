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
		Aps: IosPayloadAps{
			Alert: IosPayloadApsAlert{
				Title:    "标题",
				Subtitle: "子标题",
				Body:     "内容",
			},
		},
	}
	param := &SendParam{
		Types:          TypeUnicast,
		DeviceTokens:   config.DeviceTokens,
		AliasType:      "",
		Alias:          "",
		FileId:         "",
		Filter:         "",
		Payload:        payload,
		Policy:         Policy{},
		ProductionMode: true,
		Description:    "测试推送",
		Mipush:         true,
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
