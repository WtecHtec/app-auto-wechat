package model

import "time"

type User struct {
	Id         string    `xorm:"varchar(64) pk notnull unique 'id' comment('用户id')"`
	NickName   string    `xorm:"varchar(255)  unique 'nick_name' comment('小程序获取的用户名')" json:"nick_name"`
	PbOpenId   string    `xorm:"varchar(255)  unique 'pb_openid' comment('公众号用户openid')" json:"pb_openid"`
	ApOpenId   string    `xorm:"varchar(255)  unique 'ap_openid' comment('小程序用户openid')" json:"ap_openid"`
	Enable     bool      `xorm:"Bool notnull  'enable' default 1 comment('是否可用')"`
	CreateTime time.Time `xorm:"DateTime notnull created  'create_time' comment('创建时间')"`
	UpdateTime time.Time `xorm:"DateTime notnull updated  'update_time' comment('更新时间')"`
	OpenId     string    `json:"openid"`
}
