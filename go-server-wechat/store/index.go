package store

import (
	"serverwechat/model"

	"github.com/eatmoreapple/openwechat"
)

type WxBot struct {
	Id   string
	Bot  *openwechat.Bot
	Auto *model.Auto
}

var BOTS = make(map[string]*WxBot)
