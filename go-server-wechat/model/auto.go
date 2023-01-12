package model

import "time"

type Auto struct {
	Id         string    `xorm:"varchar(64) pk notnull unique 'id' comment('id')"`
	UserId     string    `xorm:"varchar(255) notnull  unique 'user_id' comment('用户id')" json:"user_id"`
	ReName     string    `xorm:"varchar(255)  unique 'rename' comment('备注')" json:"rename"`
	Enable     bool      `xorm:"Bool notnull  'enable' default 1 comment('是否在运行')"`
	CreateTime time.Time `xorm:"DateTime notnull created  'create_time' comment('创建时间')"`
	UpdateTime time.Time `xorm:"DateTime notnull updated  'update_time' comment('更新时间')"`
}
