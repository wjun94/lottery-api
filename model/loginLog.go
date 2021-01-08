package model

// LoginLog 用户登入表数据类型
type LoginLog struct {
	ID       string `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID   string `json:"userId"`                   // 用户id
	CreateAt string `gorm:"not null" json:"createAt"` // 登入时间
	Address  string `gorm:"not null" json:"addr"`     // 登入地点
	IP       string `gorm:"not null" json:"ip"`       // 登入的ip
	Platform string `gorm:"not null" json:"platform"` // 登入设备
	Browser  string `gorm:"not null" json:"browser"`  // 登入设备
}
