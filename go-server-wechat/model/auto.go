package model

import "time"

type Auto struct {
	Id             string    `xorm:"varchar(64) pk notnull unique 'id' comment('id')"`
	UserId         string    `xorm:"varchar(255) notnull  unique 'user_id' comment('用户id')" json:"user_id"`
	ReName         string    `xorm:"varchar(255)  'rename' comment('备注')" json:"rename"`
	AutoReply      bool      `xorm:"Bool   'auto_reply' default 0 comment('是否开启自动回复')" json:"auto_reply"`
	AutoReplyGroup bool      `xorm:"Bool   'auto_reply_group' default 0 comment('是否开启群艾特自动回复')" json:"auto_reply_group"`
	AutoBot        string    `xorm:"varchar(255)   'auto_bot' default('nobot') comment('机器人类型, tuling chatgpt')" json:"auto_bot"`
	AutoDesc       string    `xorm:"varchar(255)   'auto_desc' default('正在忙') comment('自动回复文案')" json:"auto_desc"`
	TulingApiKey   string    `xorm:"varchar(255)   'tuling_api_key' comment('图灵机器人 key')" json:"tuling_api_key"`
	Enable         bool      `xorm:"Bool  'enable' default 0 comment('是否在运行')"`
	CreateTime     time.Time `xorm:"DateTime notnull created  'create_time' comment('创建时间')"`
	UpdateTime     time.Time `xorm:"DateTime notnull updated  'update_time' comment('更新时间')"`
}
