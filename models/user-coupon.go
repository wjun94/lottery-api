package models

/**
 * 文档作者: wjun94
 * 创建时间：2019年11月05日
 * 修改时间：2019年11月05日
 * 描述信息：用户对应的优惠券数据操作
 */

import (
	"time"

	"mango-api/response"
	"strconv"

	"github.com/google/uuid"
)

// UserCoupon 用户领取的优惠券数据
type UserCoupon struct {
	ID             string  `gorm:"primary_key" json:"id"`                                      // 地址唯一标识
	UserID         string  `gorm:"not null" json:"userId"`                                     // 用户唯一标识
	CommodityID    string  `gorm:"not null" json:"commodityId,omitempty"`                      // 商品id
	Name           string  `gorm:"not null" json:"name"`                                       // 优惠券名
	Describe       string  `gorm:"not null" json:"describe"`                                   // 描述
	CreateTime     string  `gorm:"not null" json:"createTime,omitempty"`                       // 创建时间
	Unit           string  `gorm:"not null" json:"unit,omitempty"`                             // 单位
	Type           byte    `gorm:"not null" json:"type,omitempty"`                             // 类型
	Value          float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"value,omitempty"`  // 面值
	Full           float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"full,omitempty"`   // 满
	Reduce         float64 `gorm:"not null" gorm:"type:decimal(10,2)" json:"reduce,omitempty"` // 减
	ValidityPeriod int     `gorm:"not null" json:"validityPeriod,omitempty"`                   // 有效期
}

// StrToTime string转毫秒
// stringTime 日期
func StrToTime(stringTime string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		unixTime := theTime.Unix() //1504082441
		return unixTime
	}
	return 0
}

// TimeAddDay 日期添加天数
// stringTime 日期
// day 天数
func TimeAddDay(stringTime string, day int) (string, error) {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		dd, _ := time.ParseDuration(strconv.Itoa(24*day) + "h")
		dd1 := theTime.Add(dd)
		return dd1.Format("2006-01-02 15:04:05"), nil
	}
	return "", err
}

// InsertUserCoupon 用户插入优惠券
func InsertUserCoupon(userCoupon UserCoupon) response.Response {
	var result response.Response

	var coup Coupon
	Db().Where("id = ?", userCoupon.ID).Order("create_time desc").Find(&coup)
	coup.NumberSheets++
	coup.Receive++
	Db().Model(&Coupon{}).Where("id = ?", userCoupon.ID).Updates(&coup)

	userCoupon.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	userCoupon.ID = uuid.New().String()
	if dbc := Db().Create(&userCoupon); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(userCoupon)
	}
	return result
}

// FindAllUserCoupon 查询用户所有的优惠券
func FindAllUserCoupon(userID string) response.Response {
	type userCoupon struct {
		UserCoupon
	}
	var (
		data   []userCoupon
		result response.Response
	)
	nowTime := time.Now().AddDate(0, 0, -120)
	Db().Where("create_time < ?", nowTime).Delete(&userCoupon{})
	if dbc := Db().Where("user_id = ?", userID).Order("create_time desc").Find(&data); dbc.Error != nil {
		// 返回错误信息
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(data)
	}
	return result
}

// DeleteUserCouponByID 删除优惠券
func DeleteUserCouponByID(id string) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&UserCoupon{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}
