package wsqr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serverwechat/dao"
	"serverwechat/wxbot"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func MyWs(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &Connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	H.r <- c
	go c.writer()
	c.reader()
	defer func() {
		c.data.Type = "logout"
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		H.b <- data_b
		H.r <- c
	}()
}

func (c *Connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

func SendMsg(data Data) {
	data_b, _ := json.Marshal(data)
	H.b <- data_b
}

func (c *Connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			H.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		fmt.Println("message===", message)
		switch c.data.Type {
		case "login":
			_, status := dao.GetUserInfoByPbOpenId("", "", c.data.Content, "")
			c.data.Status = 404
			if status == 1 {
				wx := wxbot.NewWxBot(c.ws, c.data.Content)
				go func() {
					wx.Login()
				}()
				c.data.Status = 200
			}
			data_b, _ := json.Marshal(c.data)
			H.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			H.b <- data_b
		default:
			fmt.Print("========default================")
		}
	}
}
