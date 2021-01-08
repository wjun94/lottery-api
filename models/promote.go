package models

import (
	"mango-api/response"
	"github.com/google/uuid"
	"time"
)

// Promote 优惠券接口数据
type Promote struct {
	ID         string `gorm:"primary_key" json:"id"`
	LinkID     string `gorm:"not null" json:"linkId"`
	CouponID   string `gorm:"not null" json:"couponId"`
	Name       string `gorm:"not null" json:"name"`
	CreateTime string `json:"createTime"`                           // 创建时间
	Status     *bool  `gorm:"not null;default:false" json:"status"` // 上下架
	Img        string `gorm:"not null" json:"img"`
	Link       Link   `json:"link"`
	Coupon     Coupon `json:"coupon"`
}

// InsertPromote 创建订单
func InsertPromote(promote Promote) response.Response {
	promote.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	promote.ID = uuid.New().String()
	var result response.Response
	if dbc := Db().Create(&promote); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindPromote 获取分页弹窗和账户
func FindPromote(current int, pageSize int) response.Response {
	var promote []Promote
	var count int
	Db().Limit(pageSize).Offset((current - 1) * pageSize).Find(&promote)
	for key, v := range promote {
		Db().Model(&v).Related(&v.Link).Related(&v.Coupon)
		promote[key] = v
	}
	Db().Model(&Promote{}).Count(&count)
	return response.PageJSON(promote, count)
}

// FindPromoteByLinkID 查找对应用户(LinkID)的上架弹窗
func FindPromoteByLinkID(LinkID string) response.Response {
	var promote []Promote
	var result response.Response
	if dbc := Db().Where("link_id = ? AND status = ?", LinkID, true).Find(&promote); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(promote)
	}
	return result
}

// UpdatePromote 更新弹窗数据
func UpdatePromote(promote Promote) response.Response {
	var result response.Response
	if dbc := Db().Model(&Promote{}).Where("id = ?", promote.ID).Update(&promote); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// DeletePromote 删除弹窗数据
func DeletePromote(id string) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&Promote{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}
