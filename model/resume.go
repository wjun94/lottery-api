package model

// 用户简历表
type Resume struct {
	ID         string `gorm:"primary_key" json:"id"`
	UserID     string `json:"userID"`     // 用户id
	Name       string `json:"name"`       // 职位名称
	Content    string `json:"content"`    // 工作内容
	StartAt    string `json:"createAt"`   // 开始时间
	EndAt      string `json:"endAt"`      // 结束时间
	Department string `json:"department"` // 所属部门
	KPI        string `json:"kpi"`        // 工作业绩
}
