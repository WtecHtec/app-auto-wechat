package wsqr

import "encoding/json"

var H = hub{
	c: make(map[*Connection]bool),
	u: make(chan *Connection),
	b: make(chan []byte),
	r: make(chan *Connection),
}

type hub struct {
	c map[*Connection]bool
	b chan []byte // 消息内容
	r chan *Connection
	u chan *Connection
}

func (h *hub) Run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
