package umeng_push

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 签名生成规则:
// 提取请求方法method（POST,全大写）;
// 提取请求url信息,包括Host字段的域名(或ip:端口)和URI的path部分,注意不包括path的querystring,比如http://msg.umeng.com/api/send 或者 http://msg.umeng.com/api/status;
// 提取请求的post-body;
// 拼接请求方法、url、post-body及应用的app_master_secret;
// 将形成字符串计算MD5值,形成一个32位的十六进制（字母小写）字符串,即为本次请求sign（签名）的值;Sign=MD5($http_method$url$post-body$app_master_secret);
func Sign(url, postBody, AppMasterSecret string) string {
	s := "POST" + url + postBody + AppMasterSecret
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UrlSign(url, postBody, AppMasterSecret string) string {
	sign := Sign(url, postBody, AppMasterSecret)
	return fmt.Sprintf("%s?sign=%s", url, sign)
}
