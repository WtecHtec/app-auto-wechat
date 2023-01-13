package model

import "time"

type User struct {
	Id         string    `xorm:"varchar(64) pk notnull unique 'id' comment('用户id')"`
	WxUnid     string    `xorm:"varchar(255)  'wx_unid' comment('微信登录标识')" json:"wx_unid"`
	PbOpenId   string    `xorm:"varchar(255)  'pb_openid' comment('公众号用户openid')" json:"pb_openid"`
	ApOpenId   string    `xorm:"varchar(255)  'ap_openid' comment('小程序用户openid')" json:"ap_openid"`
	Enable     bool      `xorm:"Bool notnull  'enable' default 1 comment('是否可用')"`
	CreateTime time.Time `xorm:"DateTime notnull created  'create_time' comment('创建时间')"`
	UpdateTime time.Time `xorm:"DateTime notnull updated  'update_time' comment('更新时间')"`
	OpenId     string    `json:"openid"`
}
