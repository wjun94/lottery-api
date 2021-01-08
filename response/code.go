package response

// ErrorCode 错误码
var ErrorCode = map[uint16]string{
	200:  "success",
	304:  "请重新登入",                                   // 没有权限
	501:  "请填写完整",                                   // 缺少字段
	204:  "没有查询到数据",                                 // 没有查询到数据
	505:  "账号或密码错误",                                 // 没有查询到数据
	1062: "已注册",                                     // 已注册
	1364: "Field 'id' doesn't have a default value", // 缺少id
}
