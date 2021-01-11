package model

// Publicity 一房一价、开发商等信息表
type Publicity struct {
	ID      string `form:"id" gorm:"primary_key" json:"id"`
	LotID   string `form:"lotId" json:"lotId" label:"摇号信息表id"`
	Comp    string `form:"comp" json:"comp" label:"开发公司"`
	Name    string `form:"name" json:"name" label:"项目名称"`
	Addr    string `form:"addr" json:"addr" label:"房屋坐落位置"`
	Effect  string `form:"effect" json:"effect" label:"用途"`
	Cert    string `form:"cert" json:"cert" label:"证号"`
	PubAt   string `form:"pubAt" json:"pubAt" label:"公示日期"`
	LegalAt string `form:"legalAt" json:"legalAt" label:"核发日期"`
	Tot     string `form:"tot" json:"tot" label:"可售套数"`
	Price   string `form:"price" json:"price" label:"一房一价"`
}
