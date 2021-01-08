package model

// Recruit 招聘表类型
type Recruit struct {
	ID       string `gorm:"primary_key" json:"id"`
	UserID   string `json:"userId" label:"用户id"`
	Name     string `json:"name" label:"岗位名称"`
	Type     string `json:"type" binding:"required" label:""`
	Job      string `json:"job" binding:"required" label:"岗位类型"`
	Status   byte   `json:"status" label:"岗位状态"` // (1:未发布，2:已发布，3:暂停，4:到期)
	CreateAt string `json:"createAt" label:"创建时间"`
	UpdateAt string `json:"updateAt" label:"更新时间"`
	Views    uint32 `json:"views" label:"访问量"`
	Nat      *bool  `json:"nat" label:"工作性质"` // (全职/兼职)
	Ugt      *bool  `json:"ugt" label:"急聘"`
	Benf     string `json:"benf" label:"福利待遇"`
	Educ     byte   `json:"educ" label:"学历要求"`
	Exp      byte   `json:"exp" label:"工作经验"`
	Sex      byte   `json:"sex" label:"性别"`
	Rcrt     byte   `json:"rcrt" label:"招聘人数"`
	Prov     string `json:"prov" label:"省"`
	City     string `json:"city" label:"市"`
	Area     string `json:"area" label:"区"`
	Desc     string `json:"desc" gorm:"type:longtext" label:"描述"`
	Cont     string `json:"cont" label:"联系人"`
	Phone    string `json:"phone" label:"联系电话"`
	Email    string `json:"email" label:"联系邮箱"`
	Losal    int    `json:"losal" label:"最低薪资"`
	Hisal    int    `json:"hisal" label:"最高薪资"`
}

// ListRcrt 列表数据
type ListRcrt struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	UpdateAt string `json:"updateAt"`
	Status   byte   `json:"status"`
	Views    uint32 `json:"views"`
	Ugt      *bool  `json:"ugt"`
	Benf     string `json:"benf"`
}

type cInfo struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name" label:"公司名"`
	Ind    string `json:"ind"`
	Scale  byte   `json:"scale" label:"公司规模"`
}

type user struct {
	ID    string `json:"id"`
	CInfo cInfo  `json:"cInfo"` // 企业信息
}

// ListRcrtByComp 求职者网站列表(包含公司)
type ListRcrtByComp struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name"`
	UpdateAt string `json:"updateAt"`
	Status   byte   `json:"status"`
	Ugt      *bool  `json:"ugt"`
	Benf     string `json:"benf"`
	Losal    int    `json:"losal"`
	Hisal    int    `json:"hisal"`
	Educ     byte   `json:"educ"`
	Exp      byte   `json:"exp"`
	Rcrt     byte   `json:"rcrt"`
	User     user   `json:"user"`
}

type USER struct {
	ID    string `json:"id"`
	CInfo CInfo  `json:"cInfo"`
}

// RcrtDetail 求职者网站列表(包含公司)
type RcrtDetail struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Name     string `json:"name" label:"岗位名称"`
	Type     string `json:"type" binding:"required" label:""`
	Job      string `json:"job" binding:"required" label:"岗位类型"`
	Status   byte   `json:"status" label:"岗位状态"` // (1:未发布，2:已发布，3:暂停，4:到期)
	CreateAt string `json:"createAt" label:"创建时间"`
	UpdateAt string `json:"updateAt" label:"更新时间"`
	Views    uint32 `json:"views" label:"访问量"`
	Nat      *bool  `json:"nat" label:"工作性质"` // (全职/兼职)
	Ugt      *bool  `json:"ugt" label:"急聘"`
	Benf     string `json:"benf" label:"福利待遇"`
	Educ     byte   `json:"educ" label:"学历要求"`
	Exp      byte   `json:"exp" label:"工作经验"`
	Sex      byte   `json:"sex" label:"性别"`
	Rcrt     byte   `json:"rcrt" label:"招聘人数"`
	Prov     string `json:"prov" label:"省"`
	City     string `json:"city" label:"市"`
	Area     string `json:"area" label:"区"`
	Desc     string `json:"desc" gorm:"type:longtext" label:"描述"`
	Cont     string `json:"cont" label:"联系人"`
	Phone    string `json:"phone" label:"联系电话"`
	Email    string `json:"email" label:"联系邮箱"`
	Losal    int    `json:"losal" label:"最低薪资"`
	Hisal    int    `json:"hisal" label:"最高薪资"`
	User     USER   `json:"user"`
}

// UpdateRcrt 更新职位
type UpdateRcrt struct {
	ID       string `validate:"required" json:"id"`
	Name     string `json:"name" label:"岗位名称"`
	Type     string `json:"type" label:""`
	Job      string `json:"job" label:"岗位类型"`
	Status   byte   `json:"status" label:"岗位状态"` // (1:未发布，2:已发布，3:暂停，4:到期)
	CreateAt string `json:"createAt" label:"创建时间"`
	UpdateAt string `json:"updateAt" label:"更新时间"`
	Views    uint32 `json:"views" label:"访问量"`
	Nat      *bool  `json:"nat" label:"工作性质"` // (全职/兼职)
	Ugt      *bool  `json:"ugt" label:"急聘"`
	Benf     string `json:"benf" label:"福利待遇"`
	Educ     byte   `json:"educ" label:"学历要求"`
	Exp      byte   `json:"exp" label:"工作经验"`
	Sex      byte   `json:"sex" label:"性别"`
	Rcrt     byte   `json:"rcrt" label:"招聘人数"`
	Prov     string `json:"prov" label:"省"`
	City     string `json:"city" label:"市"`
	Area     string `json:"area" label:"区"`
	Desc     string `json:"desc" label:"描述"`
	Cont     string `json:"cont" label:"联系人"`
	Phone    string `json:"phone" label:"联系电话"`
	Email    string `json:"email" label:"联系邮箱"`
	Losal    int    `json:"losal" label:"最低薪资"`
	Hisal    int    `json:"hisal" label:"最高薪资"`
}

// UpdateRcrts 更新职位
type UpdateRcrts struct {
	Ids   []string `validate:"required" form:"ids"`
	Views *uint32  `form:"views" label:"访问量"`
	Ugt   *bool    `form:"ugt" json:"ugt" label:"急聘"`
}

// SelectRcrt 职位信息数据
type SelectRcrt struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Job   string `json:"job"`
	Nat   *bool  `json:"nat"`
	Educ  byte   `json:"educ"`
	Exp   byte   `json:"exp"`
	Ugt   *bool  `json:"ugt"`
	Sex   byte   `json:"sex"`
	Rcrt  byte   `json:"rcrt"`
	Prov  string `json:"prov"`
	City  string `json:"city"`
	Area  string `json:"area"`
	Desc  string `json:"desc"`
	Benf  string `json:"benf"`
	Cont  string `json:"cont"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Losal int    `json:"losal"`
	Hisal int    `json:"hisal"`
}
