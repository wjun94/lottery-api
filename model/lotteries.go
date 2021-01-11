package model

// Lottery 摇号房子表
type Lottery struct {
	ID        string `form:"id" gorm:"primary_key" json:"id"`
	Name      string `form:"name" json:"name" label:"项目名称"`
	Status    string `form:"status" json:"status" label:"状态"`
	CreateAt  string `form:"createAt" json:"createAt" label:"创建日期"`
	LotAt     string `form:"lotAt"  json:"lotAt" label:"摇号时间"`
	RegAt     string `form:"regAt"  json:"regAt" label:"登记时间"`
	RegEndAt  string `form:"regEndAt"  json:"regEndAt" label:"登记结束时间"`
	RegAddr   string `form:"regAddr" json:"regAddr" label:"登记地址"`
	Src       string `form:"src" json:"src" label:"哪个地方爬取"`
	Detl      string `form:"detl" json:"detl" label:"公告"`
	TotalRoom int    `form:"totalRoom"  json:"totalRoom" label:"总申请人数"`
	NoAeRoom  int    `form:"noAeRoom" json:"noAeRoom" label:"无房人才申请人数"`
	NoRoom    int    `form:"noRoom" json:"noRoom" label:"有房申请人数"`
	HavaRoot  int    `form:"havaRoot" json:"havaRoot" label:"有房申请人数"`
	PdfURL    string `form:"pdfUrl" gorm:"type:longtext" json:"pdfUrl" label:"pdf地址"`
	UTyp      string `form:"uTyp" gorm:"type:longtext" json:"uTyp" label:"户型图"`
	Claim     string `form:"claim" json:"claim" label:"要求(存款证明)"`
}
