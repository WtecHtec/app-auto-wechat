export const RE_OPT = 5 // 重复操作次数

export const PAGE_STATUS = {
	loading: -1, // 加载
	empty: 0, // 数据空
	error: 1, // 异常报错
	noright: 2, // 没有操作权限，没有授权
	nonetwork: 3, // 没有网络,接口请求失败
	normal: 4, // 正常操作
	reopt: 5, // 重复操作
  sueopt: 200, // 操作成功
	opt404: 404, // 没有数据
	opt405: 405, // 已存在
	timeout: 406,
}