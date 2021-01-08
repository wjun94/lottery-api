package models

import (
	"encoding/json"
	"strings"
	"time"

	"unsafe"

	"mango-api/response"

	"github.com/google/uuid"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月25日
 * 修改时间：2019年10月25日
 * 描述信息：商品数据库操作
 */

// LessCommodity 非表关联数据
type LessCommodity struct {
	ID         string  `gorm:"primary_key" json:"id"`                                        // 地址唯一标识
	AdImg      string  `gorm:"not null;type:longtext" json:"adImg,omitempty"`                // 商品广告位图
	Img        string  `gorm:"not null;type:longtext" json:"img,omitempty"`                  // 商品主图片
	Imgs       string  `gorm:"not null;type:longtext" json:"imgs,omitempty"`                 // 商品轮播图
	DetailImgs string  `gorm:"not null;type:longtext" json:"detailImgs,omitempty"`           // 商品轮播图
	Name       string  `gorm:"not null" json:"name"`                                         // 商品名
	Describe   string  `gorm:"not null" json:"describe,omitempty"`                           // 描述
	Price      float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"price,omitempty"`    // 价格
	OldPrice   float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"oldPrice,omitempty"` // 原价
	Unit       string  `gorm:"not null" json:"unit,omitempty"`                               // 单位
	Hot        *bool   `gorm:"type:boolean;" json:"hot,omitempty"`                           // 是否火热(1:不火热/2:火热)
	CreateTime string  `gorm:"not null" json:"createTime,omitempty"`                         // 创建时间
	UpdateTime string  `gorm:"not null" json:"updateTime,omitempty"`                         // 创建时间
	Status     *bool   `gorm:"type:boolean;" json:"status"`                                  // 状态
	Options    byte    `gorm:"not null" json:"options"`                                      // 附加选项
	IsShow     *bool   `gorm:"type:boolean;" json:"isShow"`                                  // 是否在主页显示
	Related    string  `gorm:"not null" json:"related"`                                      // 关联的商品
	Code       string  `json:"code"`                                                         // 商品编码
	Style      string  `gorm:"not null" json:"style"`                                        // 款式
	Model      string  `gorm:"not null" json:"model"`                                        // 款式下面的型号
	Width      int     `json:"width"`                                                        // 图片尺寸
	Height     int     `json:"height"`                                                       // 图片尺寸
	Extend1    byte    `gorm:"not null" json:"extend1,omitempty"`                            // 扩展字段1
	Extend2    string  `gorm:"not null" json:"extend2,omitempty"`                            // 扩展字段2
	// Type       byte    `gorm:"not null" json:"type"`                               // 涉及到页面跳转
}

// FindComp (FindCommodityUserCouponByID) 查询商品数据接口
type FindComp struct {
	Commodity
	AdImg       []string         `json:"adImg,omitempty"`
	Img         []string         `json:"img,omitempty"`
	Imgs        []string         `json:"imgs,omitempty"`
	DetailImgs  []string         `json:"detailImgs,omitempty"`
	CommodityAd bool             `json:"commodityAd,omitempty"`
	Related     []*LessCommodity `json:"related,omitempty"`
}

// Commodity 商品表数据结构
type Commodity struct {
	LessCommodity
	Coupons     []Coupon     `json:"coupons"`
	UserCoupons []UserCoupon `json:"userCoupons"`
	CommodityAd CommodityAd  `json:"commodityAd"`
}

// InsertCommodity 添加数据
// @param isAd 是否是广告位
func InsertCommodity(comm Commodity, isAd bool) response.Response {
	var (
		result response.Response
	)
	comm.ID = uuid.New().String()
	var count int
	Db().Model(&Commodity{}).Count(&count)
	comm.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	comm.UpdateTime = time.Now().Format("2006-01-02 15:04:05")

	if dbc := Db().Create(&comm); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	if count == 0 || isAd {
		var commodityAd = CommodityAd{
			CommodityID: comm.ID,
		}
		Db().Exec("truncate table commodity_ads;")
		Db().Create(&commodityAd)
	}
	return result
}

// UpdateCommodity 更新数据
func UpdateCommodity(comm Commodity, isAd bool) response.Response {
	var result response.Response
	if isAd {
		var commodityAd = CommodityAd{
			CommodityID: comm.ID,
		}
		Db().Exec("truncate table commodity_ads;")
		Db().Model(&CommodityAd{}).Create(&commodityAd)
	}
	if dbc := Db().Model(&Commodity{}).Where("id = ?", comm.ID).Updates(&comm); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindAllCommodity 查询所有商品数据
func FindAllCommodity() response.Response {
	var comms []Commodity
	Db().Select("id, name").Find(&comms)
	var data []map[string]interface{}
	base, _ := json.Marshal(comms)
	json.Unmarshal(base, &data)
	return response.JSON(data)
}

// FindRelatedCommodity 所有商品和优惠券关联查询 PC端后台使用
// current 当前页码
// pageSize 页面个数
func FindRelatedCommodity(current int, pageSize int) response.Response {
	var (
		comms []Commodity
		count int
		data  []struct {
			Commodity
			Imgs        []string         `json:"imgs"`
			DetailImgs  []string         `json:"detailImgs"`
			CommodityAd bool             `json:"commodityAd"`
			Related     []*LessCommodity `json:"related"`
		}
	)
	Db().Limit(pageSize).Offset((current - 1) * pageSize).Find(&comms)
	base, _ := json.Marshal(comms)
	json.Unmarshal(base, &data)
	for i := range data {
		relatedArr := make([]*LessCommodity, 0)
		data[i].Imgs = strings.Split(comms[i].Imgs, ",")
		data[i].DetailImgs = strings.Split(comms[i].DetailImgs, ",")
		// 返回优惠券数据
		data[i].Coupons = FindCouponByCommodityID(comms[i].ID, "*")
		// 是否是广告位
		data[i].CommodityAd = FindCommodityAdById(comms[i].ID)
		// 查找关联商品
		related := strings.Split(comms[i].Related, ",")
		for _, v := range related {
			for _, child := range comms {
				if child.ID == v {
					relatedArr = append(relatedArr, (*LessCommodity)(unsafe.Pointer(&child)))
					break
				}
			}
		}
		data[i].Related = relatedArr
	}
	Db().Model(&Commodity{}).Count(&count)
	return response.PageJSON(data, count)
}

// FindCommodityUserCouponByID 查找用户商品对应的优惠券(小程序详情页获取用户优惠券和商品优惠券 经典案例，多表)
func FindCommodityUserCouponByID(id string, userID string) response.Response {
	type userCoupon struct {
		UserCoupon
	}
	type commodity struct {
		ID          string       `json:"id"`          // 地址唯一标识
		Coupons     []Coupon     `json:"coupons"`     // 优惠券
		UserCoupons []userCoupon `json:"userCoupons"` // 所有的用户优惠券
	}
	var (
		comms commodity
	)
	Db().Where("id=?", id).Find(&comms)
	Db().Model(&comms).Related(&comms.Coupons).Order("create_time desc").Where("user_id = ? AND date_add(create_time, interval validity_period || 1 day) > ?", userID, time.Now().Format("2006-01-02 15:04:05")).Related(&comms.UserCoupons)
	return response.JSON(comms)
}

// FindCommodity 查询id对应的商品
func FindCommodity(id string) FindComp {
	var (
		commodity Commodity
		data      FindComp
	)
	Db().Where("id = ?", id).First(&commodity)
	base, _ := json.Marshal(commodity)
	json.Unmarshal(base, &data)
	data.AdImg = strings.Split(commodity.AdImg, ",")
	data.Img = strings.Split(commodity.Img, ",")
	data.Imgs = strings.Split(commodity.Imgs, ",")
	data.DetailImgs = strings.Split(commodity.DetailImgs, ",")
	return data
}

// FindCommodityList 获取列表
func FindCommodityList() response.Response {
	var (
		commodities []Commodity
	)
	type result struct {
		LessCommodity
		CommodityAd bool `json:"commodityAd"`
	}
	Db().Find(&commodities)
	toResult := make([]*result, 0)
	for i, v := range commodities {
		var progress = (*result)(unsafe.Pointer(&commodities[i]))
		progress.CommodityAd = FindCommodityAdById(v.ID)
		toResult = append(toResult, progress)
	}
	return response.JSON(toResult)
}

// DeleteCommodity 删除数据
func DeleteCommodity(id string) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&Commodity{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}
