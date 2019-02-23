package e

const (
	OK                = 0   // 成功
	ErrSystemCode     = 500 //内部错误
	ErrParamCode      = 501 //输入参数错误
	ErrNoUserIdCode   = 502 //uid缺失
	ErrInvalidCode    = 503 //接口废弃
	ErrReqTimeoutCode = 504 //请求参数过期
	ErrTimeoutCode    = 505 //超时
	ErrTokenCode      = 506 //Token验证错误
	ErrSignCode       = 507 //Sign验证错误
	ErrNoServiceCode  = 508 //未找到服务

	//自定义错误码6位数字（2位业务模块+4位错误异常
	//业务模块类型占用2位 比如 user 模块为30 order 模块为 40 ...
	CustomerErrorCode = 200001 + iota
)

var CodeDescMap = map[int]string{
	OK:                "成功",
	ErrSystemCode:     "内部错误",
	ErrParamCode:      "输入参数错误",
	ErrNoUserIdCode:   "uid缺失",
	ErrInvalidCode:    "接口废弃",
	ErrReqTimeoutCode: "请求参数过期",
	ErrTimeoutCode:    "超时",
	ErrTokenCode:      "Token验证错误",
	ErrSignCode:       "Sign验证错误",
	ErrNoServiceCode:  "未找到服务",

	// 自定义
	CustomerErrorCode: "自定义异常",
}
