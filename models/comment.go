package models

import (
	"mango-api/response"
	"time"
)

// Comment 评论数据
type Comment struct {
	OrderID     string `gorm:"primary_key" json:"orderId"`           // 订单号
	UserID      string `json:"userId"`                               // 用户id
	CommodityID string `json:"commodityId"`                          // 商品id
	CreateTime  string `json:"createTime"`                           // 创建时间
	Product     int    `json:"product"`                              // 产品评级
	Logistics   int    `json:"logistics"`                            // 物流评级
	Images      string `gorm:"not null;type:longtext" json:"images"` // 图片
	Detail      string `json:"detail"`                               // 描述
}

// InsertComment 添加评论
func InsertComment(comment Comment) response.Response {
	var (
		result response.Response
	)
	comment.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if dbc := Db().Create(&comment); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindComment 获取用户评论
func FindComment(orderID string) response.Response {
	var (
		result  response.Response
		comment Comment
	)
	if dbc := Db().Where("order_id = ?", orderID).Find(&comment); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(comment)
	}
	return result
}

// FindCommodityComments 获取商品对应评论
func FindCommodityComments(commodityID string, current int, pageSize int) response.Response {
	type comment struct {
		Comment
		User
	}
	var (
		comm  []comment
		count int
	)
	if dbc := Db().Order("create_time desc").Where("commodity_id = ?", commodityID).Limit(pageSize).Offset((current - 1) * pageSize).Find(&comm); dbc.Error != nil {
		return response.JSONError(dbc.Error.Error())
	}

	for i, v := range comm {
		Db().Model(&v.Comment).Related(&v.User)
		comm[i] = v
	}
	Db().Model(&Comment{}).Count(&count)
	return response.PageJSON(comm, count)
}
