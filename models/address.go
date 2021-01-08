package models

import (
	"mango-api/response"
	"encoding/json"
	"github.com/google/uuid"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：用户地址数据库操作
 */

// Address 地址表数据结构
type Address struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`              // 地址唯一标识
	UserID   string    `gorm:"not null" json:"userId,omitempty"`   // 用户id
	Name     string    `gorm:"not null" json:"name,omitempty"`     // 头像链接
	Province string    `gorm:"not null" json:"province,omitempty"` // 省
	City     string    `gorm:"not null" json:"city,omitempty"`     // 城市
	Area     string    `gorm:"not null" json:"area,omitempty"`     // 区
	Address  string    `gorm:"not null" json:"address,omitempty"`  // 详细地址
	Phone    string    `gorm:"not null" json:"phone,omitempty"`    // 手机
	User     User      `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID" json:"user,omitempty"`
}

// VerifyAddress 输入校验
type VerifyAddress struct {
	UserID   string `validate:"required" label:"用户id"`
	Name     string `validate:"required" label:"姓名"`
	Province string `validate:"required" label:"省"`
	City     string `validate:"required" label:"城市"`
	Address  string `validate:"required" label:"详细地址"`
	Phone    string `validate:"required,min=6,max=11,numeric" label:"手机号"`
}

// InsertAddress 添加数据
// address 地址数据
// isDefault 是否是默认地址
func InsertAddress(address Address, isDefault bool) response.Response {
	var (
		result response.Response
	)
	address.ID = uuid.New()
	if isDefault {
		// 更新用户表
		user := User{
			AddressID: address.ID.String(),
		}
		Db().Model(&User{}).Where("id = ?", address.UserID).Update(&user)
	}
	if dbc := Db().Create(&address); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// UpdateAddress 更新数据
// address 地址数据
// isDefault 是否是默认地址
func UpdateAddress(address Address, isDefault bool) response.Response {
	var result response.Response
	var user User
	if isDefault {
		// 更新用户表
		user = User{
			AddressID: address.ID.String(),
		}
	} else {
		user = User{
			AddressID: "-",
		}
	}
	Db().Model(&User{}).Where("id = ?", address.UserID).Updates(&user)
	if dbc := Db().Model(&Address{}).Where("id = ?", address.ID).Updates(&address); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindAllAddress 查询所有数据
// userID 用户id
func FindAllAddress(userID string) response.Response {
	var ads []Address
	var result []map[string]interface{}
	Db().Where("user_id = ?", userID).Find(&ads)
	res, _ := json.Marshal(ads)
	json.Unmarshal(res, &result)
	for key, v := range ads {
		Db().Model(&v).Related(&v.User)
		delete(result[key], "user")
		result[key]["isDefault"] = v.User.AddressID == v.ID.String()
	}
	return response.JSON(result)
}

// DeleteAddress 删除地址数据
func DeleteAddress(id uuid.UUID) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&Address{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}
