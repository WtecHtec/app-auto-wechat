package wxbot

import (
	"encoding/json"
	"fmt"
	"serverwechat/bot"
	"serverwechat/dao"
	"serverwechat/model"
	"serverwechat/store"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
	"github.com/skip2/go-qrcode"
)

type WxBot struct {
	Bot openwechat.Bot
	ws  *websocket.Conn
	cId string
}

type data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
	Status   int      `json:"status"`
}

func NewWxBot(ws *websocket.Conn, c string) *WxBot {
	return &WxBot{ws: ws, cId: c}
}

func (w *WxBot) ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	// fmt.Println(q.ToString(true))
	qrimg := "qrcodes/" + uuid + ".png"
	q.WriteFile(256, qrimg)
	data_b, _ := json.Marshal(&data{
		Status:  0,
		Content: qrimg,
		Type:    "wxcodeqr",
		User:    w.cId,
	})
	w.ws.WriteMessage(websocket.TextMessage, data_b)
}
func (w *WxBot) Login() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.UUIDCallback = w.ConsoleQrCode
	bot.MessageHandler = func(msg *openwechat.Message) {
		fmt.Println("接收信息===：", msg.Content, store.BOTS[w.cId].Auto)
		if store.BOTS[w.cId].Auto == nil {
			return
		}
		if msg.IsSendByFriend() && store.BOTS[w.cId].Auto.AutoReply == true {
			replyText(msg, *store.BOTS[w.cId].Auto)
		}
		if msg.IsAt() && store.BOTS[w.cId].Auto.AutoReply == true && store.BOTS[w.cId].Auto.AutoReplyGroup == true {
			replyText(msg, *store.BOTS[w.cId].Auto)
		}
	}

	bot.ScanCallBack = func(body []byte) {
		data_b, _ := json.Marshal(&data{
			Status:  1,
			Content: "扫码成功",
			Type:    "wxcodeqr",
			User:    w.cId,
		})
		w.ws.WriteMessage(websocket.TextMessage, data_b)
	}

	bot.LoginCallBack = func(body []byte) {
		res := string(body)
		if strings.Contains(res, "200") {
			// self, err := bot.GetCurrentUser()
			// if err != nil {
			// 	data_b, _ := json.Marshal(&data{
			// 		Status:  -1,
			// 		Content: "登录失败",
			// 		Type:    "wxcodeqr",
			// 		User:    w.cId,
			// 	})
			// 	w.ws.WriteMessage(websocket.TextMessage, data_b)
			// 	return
			// }
			// self.ID()
			data_b, _ := json.Marshal(&data{
				Status:  2,
				Content: "登录成功",
				Type:    "wxcodeqr",
				User:    w.cId,
			})
			// 绑定全局
			// fmt.Println("store.BOTS[w.cId]==", store.BOTS[w.cId])
			// if store.BOTS[w.cId] == nil {
			// 	store.BOTS[w.cId] = &store.WxBot{
			// 		Id:  w.cId,
			// 		Bot: &w.Bot,
			// 	}
			// } else {
			// 	// fmt.Println("store.BOTS[w.cId]== 1111")
			// 	store.BOTS[w.cId].Bot = &w.Bot
			// }
			dao.UpdateAutoEnable(w.cId, true)
			w.ws.WriteMessage(websocket.TextMessage, data_b)
		} else {
			data_b, _ := json.Marshal(&data{
				Status:  -1,
				Content: "登录失败",
				Type:    "wxcodeqr",
				User:    w.cId,
			})
			w.ws.WriteMessage(websocket.TextMessage, data_b)
		}
	}
	bot.Login()
	if store.BOTS[w.cId] == nil {
		store.BOTS[w.cId] = &store.WxBot{
			Id:  w.cId,
			Bot: bot,
		}
	} else {
		// fmt.Println("store.BOTS[w.cId]== 1111")
		store.BOTS[w.cId].Bot = bot
	}
	w.Bot = *bot
	bot.Block()
}

func replyText(msg *openwechat.Message, auto model.Auto) {
	if auto.AutoBot == "tuling" && auto.TulingApiKey != "" {
		msgs := strings.Split(msg.Content, " ")
		info := msg.Content
		if msg.IsAt() && len(msgs) > 1 {
			info = msgs[1]
		}
		msg.ReplyText(bot.TlBot(info, auto.TulingApiKey))
		return
	}
	if auto.AutoBot == "chatgpt" {
		return
	}
	if auto.AutoDesc != "" {
		msg.ReplyText(auto.AutoDesc)
	}
}
