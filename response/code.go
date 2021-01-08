package response

// ErrorCode 错误码
var ErrorCode = map[int]string{
	200:  "success",
	304:  "请先登入",
	500:  "服务器异常，请稍后重试",
	504:  "图片上传失败",
	1001: "缺少参数",                              // 缺少参数
	1002: "mobile number has been registered", // 手机号已被注册
	1003: "mailbox has been registered",       // 邮箱已被注册
	1004: "username has been registered",      // 用户名已被注册
	1005: "password is empty",                 // 用户名已被注册
}
