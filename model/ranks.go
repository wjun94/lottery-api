package model

// Rank 摇号表
type Rank struct {
	ID     string `form:"id" gorm:"primary_key" json:"id"`
	LotID  string `form:"lotId" json:"lotId" label:"摇号表id"`
	Rk     string `form:"rk" json:"rk" label:"摇号排名"`
	LotNum string `form:"lotNum" json:"lotNum" label:"摇号编号"`
}
