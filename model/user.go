package model

// User 用户表
type User struct {
	ID       string `form:"id" gorm:"primary_key" json:"id"`
	Phone    string `form:"phone" json:"phone"`                  // 注册的联系电话
	Pwd      string `form:"pwd" json:"pwd"`                      // 密码
	CreateAt string `form:"createAt" json:"createAt"`            // 创建时间
	UpdateAt string `form:"updateAt" json:"updateAt"`            // 创建时间
	Level    byte   `form:"level" gorm:"DEFAULT:3" json:"level"` // 用户等级(1:老板 2:业务员 3:企业用户 4:求职用户)
	Email    string `form:"email" json:"email"`                  // 邮箱
	CInfo    CInfo  `json:"cInfo"`                               // 企业信息
	UInfos   UInfo  `json:"uInfos"`                              // 用户信息
}

// VerifyUserLogin 登入校验
type VerifyUserLogin struct {
	Phone string `json:"phone" binding:"required" validate:"required,min=6,max=11" label:"电话"` // 注册的联系电话
	Pwd   string `json:"pwd" binding:"required" validate:"required,max=20,min=6" label:"密码"`   // 密码
}

// VerifyUserRegist 注册
type VerifyUserRegist struct {
	Phone    string `json:"phone" binding:"required" validate:"numeric,required,min=6,max=11" label:"电话"` // 注册的联系电话
	Pwd      string `json:"pwd" binding:"required" validate:"numeric,required,max=20,min=6" label:"密码"`   // 密码
	Level    byte   `json:"level" validate:"numeric"`                                                     // 用户等级(1:超级管理员 2:管理员 3:普通用户 4:企业用户)// 密码
	CreateAt string `json:"createAt"`                                                                     // 创建时间
	UpdateAt string `json:"updateAt"`                                                                     // 创建时间
}
