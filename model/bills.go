package model

// Bills 用户账单
type Bills struct {
	ID       string  `gorm:"primary_key" json:"id"`
	UserID   string  `gorm:"not null" json:"userID"`   // 用户id
	CreateAt string  `gorm:"not null" json:"createAt"` // 创建时间
	Price    float64 `gorm:"not null" json:"price"`    // 金额
	Status   byte    `json:"status"`                   // 状态(0:未付款，1:已付款)
}
