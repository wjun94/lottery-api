package model

// Work 工作、实习经历
type Work struct {
	ID      string `gorm:"primary_key" json:"id"`
	UserID  string `json:"UserID"`
	StartAt string `validate:"required" json:"createAt" label:"开始时间"`
	EndAt   string `validate:"required" json:"endAt" label:"结束时间"`
	Comp    string `validate:"required" json:"comp" label:"公司"`
	Pos     string `validate:"required" json:"pos" label:"职位"`
	Ind     string `validate:"required" json:"ind" label:"行业"`
	Dep     string `validate:"required" json:"dep" label:"部门"`
	Desc    string `validate:"required" json:"desc" label:"工作描述"`
	Type    byte   `validate:"required" json:"type" labbel:"工作类型"`
	Scale   byte   `validate:"required" json:"scale" label:"公司规模"`
}
