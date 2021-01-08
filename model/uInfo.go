package model

// UInfo Create 求职者注册的相关信息
type UInfo struct {
	ID        string `gorm:"primary_key" json:"id"`
	UserID    string `json:"userId"`
	Resume    string `json:"resume" label:"简历名称"`
	Name      string `validate:"required" json:"name" label:"姓名"`
	Phone     string `validate:"required" json:"phone" label:"手机号"`
	Email     string `validate:"required" json:"email" label:"邮箱"`
	WorkAt    string `validate:"required" json:"workAt" label:"参加工作年份"`
	Sex       byte   `validate:"required" json:"sex" label:"性别"`
	Bday      string `validate:"required" json:"bday" label:"生日"`
	Prov      string `validate:"required" json:"prov" label:"现所在省"`
	City      string `validate:"required" json:"city" label:"现所在市"`
	Area      string `json:"area" label:"现所在区"`
	HukouProv string `validate:"required" json:"hukouProv" label:"户口所在省"`
	HukouCity string `validate:"required" json:"hukouCity" label:"户口所在市"`
	HukouArea string `json:"hukouArea" label:"户口所在区"`
	Addr      string `json:"addr" label:"详细地址"`
	Marriage  byte   `json:"marriage,omitempty" label:"婚姻状况"`
	Pol       byte   `json:"pol,omitempty" label:"政治面貌"`
	Height    byte   `json:"height,omitempty" label:"身高"`
	IDCard    string `json:"idCard,omitempty" label:"身份证号"`
	Tool      byte   `json:"tool,omitempty" label:"其他聊天方式"`
	Contact   string `json:"contact,omitempty" label:"账号"`
	Avatar    string `gorm:"type:longtext" json:"avatar" label:"头像"`
}

// UpdateUInfo Update 更新
type UpdateUInfo struct {
	ID        string `validate:"required" json:"id"`
	UserID    string `json:"userId"`
	Resume    string `json:"resume" label:"简历名称"`
	Name      string `json:"name" label:"姓名"`
	Phone     string `json:"phone" label:"手机号"`
	Email     string `json:"email" label:"邮箱"`
	WorkAt    string `json:"workAt" label:"参加工作年份"`
	Sex       byte   `json:"sex" label:"性别"`
	Bday      string `json:"bday" label:"生日"`
	Prov      string `json:"prov" label:"现所在省"`
	City      string `json:"city" label:"现所在市"`
	Area      string `json:"area" label:"现所在区"`
	HukouProv string `json:"hukouProv" label:"户口所在省"`
	HukouCity string `json:"hukouCity" label:"户口所在市"`
	HukouArea string `json:"hukouArea" label:"户口所在区"`
	Addr      string `json:"addr" label:"详细地址"`
	Marriage  byte   `json:"marriage" label:"婚姻状况"`
	Pol       byte   `json:"pol" label:"政治面貌"`
	Height    byte   `json:"height" label:"身高"`
	IDCard    string `json:"idCard" label:"身份证号"`
	Tool      byte   `json:"tool" label:"其他聊天方式"`
	Contact   string `json:"contact" label:"账号"`
	Avatar    string `gorm:"type:longtext" json:"avatar" label:"头像"`
}

// type UIntention struct {
// 	Losal int `json:"losal" label:"最低薪资"`
// 	Hisal int `json:"hisal" label:"最高薪资"`
// }
