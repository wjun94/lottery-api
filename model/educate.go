package model

// Educate 求职者教育经历
type Educate struct {
	ID      string `gorm:"primary_key" json:"id"`
	UserID  string `gorm:"not null" json:"userID"`  // 用户id
	School  string `gorm:"not null" json:"school"`  // 学校
	Educate string `gorm:"not null" json:"educate"` // 学历
	Spec    string `gorm:"not null" json:"spec"`    // 专业
	EndAt   string `gorm:"not null" json:"endAt"`   // 时间段
	Exp     string `gorm:"not null" json:"exp"`     // 在校经历
}
