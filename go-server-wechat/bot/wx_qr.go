package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"serverwechat/config"

	"github.com/go-resty/resty/v2"
)

const (
	KEY_API = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
	GEN_QR  = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="
)

type WxAeKey struct {
	AtToken string `json:"access_token"`
}

type Qr struct {
	Errcode int    `json:"errcode"`
	Buffer  []byte `json:"buffer"`
}

func GetWxKey() string {
	client := resty.New() // 创建一个restry客户端
	resp, err := client.R().Get(fmt.Sprintf("%v&appid=%v&secret=%v", KEY_API, config.BASE_CONFIG.WxConfig.Appid, config.BASE_CONFIG.WxConfig.AppSecret))
	if err != nil {
		fmt.Printf("请求接口失败：%v\n", err)
		return ""
	}
	var resContent WxAeKey
	fmtErr := json.Unmarshal([]byte(resp.String()), &resContent)
	if fmtErr != nil {
		fmt.Printf("解析错误: %v\n", fmtErr)
		return ""
	}
	return resContent.AtToken
}

func GencQr(attoken string, id string) {
	if attoken == "" {
		return
	}
	client := resty.New() // 创建一个restry客户端
	resp, err :=
		client.R().EnableTrace().SetBody([]byte(
			fmt.Sprintf(`{
			"path":"page/index?id=%v",
			"width":430
		 } `, id))).Post(GEN_QR + attoken)
	if err != nil {
		fmt.Printf("请求接口失败：%v\n", err)
		return
	}
	// var resContent map[string]interface{}
	// fmt.Println("resp.String()===", resp.String())
	// fmtErr := json.Unmarshal([]byte(resp.String()), &resContent)
	// if fmtErr != nil {
	// 	fmt.Printf("解析错误: %v\n", fmtErr)
	// 	return
	// }
	fmt.Println("解析成功: ", id, attoken)
	ioutil.WriteFile("public/"+id+".png", []byte(resp.String()), 0666)
}
