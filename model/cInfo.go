package model

// CInfo 公司注册信息和其他相关信息
type CInfo struct {
	ID     string `form:"id" gorm:"primary_key" json:"id"`
	UserID string `form:"userId" gorm:"not null" json:"userId"`
	Name   string `form:"name" gorm:"not null" json:"name" label:"公司名"`
	Ind    string `form:"ind" gorm:"not null" json:"ind" label:"所属行业"`
	Prov   string `form:"prov" gorm:"not null" json:"prov" label:"省"`
	City   string `form:"city" gorm:"not null" json:"city" label:"城市"`
	Area   string `form:"area" gorm:"not null" json:"area" label:"区"`
	Addr   string `form:"addr" gorm:"not null" json:"addr" label:"详细地址"`
	Scale  byte   `form:"scale" gorm:"not null" json:"scale" label:"公司规模"`
	Cont   string `form:"cont" gorm:"not null" json:"cont" label:"联系人"`
	Logo   string `form:"logo" gorm:"not null;type:longtext" json:"logo" label:"公司图标"`
	QrCode string `form:"qrCode" gorm:"not null;type:longtext" json:"qrCode" label:"公众号图片"`
	Desc   string `form:"desc" gorm:"type:longtext" json:"desc" label:"公司简介"`
	Web    string `form:"web" json:"web" label:"企业官网"`
	Level  string `form:"level" gorm:"not null" json:"level" label:"会员"` // 会员(t1:普通用户 t2：t2会员 t3:t3会员)
}

// VerUpdCoy 更新公司信息
type VerUpdCoy struct {
	UserID string
	Name   string `from:"name" validate:"required" label:"公司名"`
	Ind    string `form:"ind" validate:"required" label:"所属行业"`
	Scale  byte   `form:"scale" validate:"required" label:"公司规模"`
	Addr   string `form:"addr" validate:"required" label:"详细地址"`
	Desc   string `form:"desc" validate:"required" label:"公司简介"`
	Prov   string `form:"prov" validate:"required" label:"省"`
	City   string `form:"city" validate:"required" label:"城市"`
	Web    string `form:"web" label:"省"`
}
