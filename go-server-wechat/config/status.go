package config

const (
	STATUS_ERROR   = 201
	STATUS_RE      = 202
	STATUS_SUE     = 200
	STATUS_RUQED   = 203
	STATUS_TIMEOUT = 210
)

var STATUS_MSG = map[int]string{
	200: "数据操作成功",
	201: "数据操作失败",
	202: "数据有重复",
	203: "参数检验不通过",
	210: "数据已过期",
}
