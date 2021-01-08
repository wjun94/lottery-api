package model

// User 用户表
type User struct {
	ID       string `form:"id" gorm:"primary_key" json:"id"`
	Level    byte   `form:"level" gorm:"DEFAULT:3" json:"level"` // 用户等级(1:管理员 2:推广商)
	CreateAt string `form:"createAt" json:"createAt"`            // 创建时间
	UpdateAt string `form:"updateAt" json:"updateAt"`            // 创建时间
	Phone    string `form:"phone" json:"phone"`                  // 注册的联系电话
	Pwd      string `form:"pwd" json:"pwd"`                      // 密码
}

// VerifyLogin 校验登陆
type VerifyLogin struct {
	Phone string `json:"phone" binding:"required" validate:"required,min=6,max=11" label:"账号"` // 注册的联系电话
	Pwd   string `json:"pwd" binding:"required" validate:"required,max=20,min=6" label:"密码"`   // 密码
}
