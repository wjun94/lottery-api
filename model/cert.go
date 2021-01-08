package model

// Educate 求职者相关证书
type Cert struct {
	ID     string `gorm:"primary_key" json:"id"`
	UserID string `gorm:"not null" json:"userID"` // 用户id
	Name   string `gorm:"not null" json:"name"`   // 证书名
	No     string `gorm:"not null" json:"no"`     // 证书号
}
