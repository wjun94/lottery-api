package model

// Intention 简历-求职意向
type Intention struct {
	ID     string `json:"id" gorm:"primary_key" label:"用户id`
	Prov   string `validate:"required" json:"prov" label:"工作所在省"`
	City   string `validate:"required" json:"city" label:"工作所在市"`
	Area   string `json:"area" label:"工作所在区"`
	Pos    string `validate:"required" json:"pos" label:"职位"`
	Ind    string `validate:"required" json:"ind" label:"行业"`
	Eval   string `validate:"required" json:"eval" label:"自我评价"`
	ComeAt byte   `json:"comeAt" label:"到岗时间"`
	Type   byte   `json:"type" labbel:"工作类型"`
	Losal  int    `json:"losal" validate:"required" label:"最低薪资"`
	Hisal  int    `json:"hisal" validate:"required" label:"最高薪资"`
}

// UIntention 更新
type UIntention struct {
	ID     string `validate:"required" json:"id" gorm:"primary_key" label:"用户id`
	Prov   string `json:"prov" label:"工作所在省"`
	City   string `json:"city" label:"工作所在市"`
	Area   string `json:"area" label:"工作所在区"`
	Pos    string `json:"pos" label:"职位"`
	Ind    string `json:"ind" label:"行业"`
	Eval   string `json:"eval" label:"自我评价"`
	ComeAt byte   `json:"comeAt" label:"到岗时间"`
	Type   byte   `json:"type" labbel:"工作类型"`
	Losal  int    `json:"losal" label:"最低薪资"`
	Hisal  int    `json:"hisal" label:"最高薪资"`
}
