package models

import (
	"mango-api/response"

	"github.com/google/uuid"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月29日
 * 修改时间：2019年10月29日
 * 描述信息：优惠券数据库操作
 */

// Coupon 优惠券表数据结构
type Coupon struct {
	ID             string    `gorm:"primary_key" json:"id"`                                      // 地址唯一标识
	CommodityID    string    `gorm:"not null" json:"commodityId,omitempty"`                      // 商品id
	Name           string    `gorm:"not null" json:"name"`                                       // 优惠券名
	Describe       string    `gorm:"not null" json:"describe"`                                   // 描述
	CreateTime     string    `gorm:"not null" json:"createTime,omitempty"`                       // 创建时间
	InvalidTime    string    `gorm:"not null" json:"invalidTime,omitempty"`                      // 优惠券结束日期
	Unit           string    `gorm:"not null" json:"unit,omitempty"`                             // 单位
	Type           byte      `gorm:"not null" json:"type,omitempty"`                             // 类型
	Value          float64   `gorm:"not null" gorm:"type:decimal(10,2)" json:"value,omitempty"`  // 面值
	Full           float64   `gorm:"not null" gorm:"type:decimal(10,2)" json:"full,omitempty"`   // 满
	Reduce         float64   `gorm:"not null" gorm:"type:decimal(10,2)" json:"reduce,omitempty"` // 减
	Amount         int       `gorm:"not null;default:-1" json:"amount,omitempty"`                // 总张数
	NumberSheets   int       `gorm:"not null" json:"numberSheets,omitempty"`                     // 领取张数
	Receive        int       `gorm:"not null" json:"receive,omitempty"`                          // 领取人数
	ValidityPeriod int       `gorm:"not null" json:"validityPeriod,omitempty"`                   // 有效期
	Count          int       `gorm:"not null" json:"count,omitempty"`                            // 发放数量
	Status         *bool     `gorm:"not null" json:"status,omitempty"`                           // 状态(上下架)
	Commodity      Commodity `gorm:"FOREIGNKEY:CommodityID;ASSOCIATION_FOREIGNKEY:ID" json:"commodity,omitempty"`
}

// InsertCoupon 添加数据
func InsertCoupon(coupon Coupon) response.Response {
	var (
		result response.Response
	)
	coupon.ID = uuid.New().String()
	if dbc := Db().Create(&coupon); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindAllCoupon 查询所有数据
func FindAllCoupon(current int, pageSize int) response.Response {
	var coup []Coupon
	var count int
	Db().Limit(pageSize).Offset((current - 1) * pageSize).Find(&coup)
	for i, v := range coup {
		Db().Model(&v).Related(&v.Commodity)
		coup[i] = v
	}

	Db().Model(&Coupon{}).Count(&count)
	return response.PageJSON(coup, count)
}

// FindCouponByCommodityID 根据商品id查询相关优惠券和商品
func FindCouponByCommodityID(id string, field string) []Coupon {
	var coupById []Coupon
	// 关联查询
	if dbc := Db().Select(field).Where("commodity_id = ?", id).Find(&coupById); dbc.Error != nil {
		// 返回错误信息
		return nil
	}
	return coupById
}

// FindCouponByID 根据id查询优惠券和相关商品
func FindCouponByID(id string) response.Response {
	var coup Coupon
	// 关联查询
	if dbc := Db().Where("id = ?", id).Find(&coup); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	Db().Model(&coup).Related(&coup.Commodity)
	return response.JSON(coup)
}

// DeleteCoupon 删除数据
func DeleteCoupon(id string) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&Coupon{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}

// UpdateCoupon 更新数据
func UpdateCoupon(coupon Coupon) response.Response {
	var result response.Response
	if dbc := Db().Model(&Coupon{}).Where("id = ?", coupon.ID).Updates(&coupon); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}
