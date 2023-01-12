package wxbot

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

type WxBot struct {
	Bot *openwechat.Bot
}

func NewWxBot() *WxBot {
	return &WxBot{}
}

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
	// q.WriteFile(256, "qrcodes/"+uuid+".png")
}

func (w *WxBot) Login() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	w.Bot = bot
	bot.UUIDCallback = ConsoleQrCode
	bot.MessageHandler = func(msg *openwechat.Message) {
		fmt.Println("接收信息===：", msg.Content)
	}
	bot.LoginCallBack = func(body []byte) {
		res := string(body)
		fmt.Println("登录状态===:", res)
	}
	bot.Login()
	bot.Block()
}
