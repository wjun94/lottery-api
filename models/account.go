package models

import (
	"mango-api/db"
	"mango-api/response"
	"mango-api/utils"
	"time"

	"github.com/google/uuid"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年11月01日
 * 修改时间：2019年11月01日
 * 描述信息：账户数据库操作
 */

// Account 用户表数据结构
type Account struct {
	ID         string `gorm:"primary_key" json:"id"`
	Name       string `gorm:"not null" json:"name"`       // 用户名
	Password   string `gorm:"not null" json:"password"`   // 密码
	Phone      string `gorm:"not null" json:"phone"`      // 手机号
	Type       byte   `gorm:"not null" json:"type"`       // 账户类型(1渠道商/2超级管理员/3卖家账号) 级别
	CreateTime string `gorm:"not null" json:"createTime"` // 创建时间
	Links      []Link `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:AccountID" json:"links"`
}

// AdminJwtToken 后台用户登入返回数据
type AdminJwtToken struct {
	Token string `json:"token"`
	Type  byte   `gorm:"not null" json:"type"` // 层级
}

// InsertAccount 添加数据
func InsertAccount(account Account) response.Response {
	var (
		result response.Response
	)
	account.ID = uuid.New().String()
	account.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if dbc := Db().Create(&account); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// FindAllAccount 用户和优惠券关联查询
func FindAllAccount(current int, pageSize int, id string) response.Response {
	var account []Account
	var count int
	Db().Limit(pageSize).Offset((current-1)*pageSize).Where("id = ?", id).Find(&account)
	if account[0].Type == 2 {
		Db().Limit(pageSize).Offset((current - 1) * pageSize).Find(&account)
	}
	for key, v := range account {
		Db().Model(&v).Related(&v.Links)
		account[key] = v
	}
	Db().Model(&Account{}).Count(&count)
	return response.PageJSON(account, count)
}

// UpdateAccount 更新数据
func UpdateAccount(account Account) response.Response {
	var result response.Response
	if dbc := Db().Model(&Account{}).Where("id = ?", account.ID).Update(&account); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}

// GetAccountByID 根据id查询用户信息
func GetAccountByID(id string) response.Response {
	var result response.Response
	var account Account
	if dbc := Db().Where("id = ?", id).Find(&account); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		Db().Model(&account).Related(&account.Links)
		result = response.JSON(account)
	}
	return result
}

// DeleteAccount 删除数据
func DeleteAccount(id string) response.Response {
	var links []Link
	Db().Select("*").Find(&links)
	for range links {
		Db().Where("account_id = ?", id).Delete(&Link{})
	}
	if dbc := Db().Where("id = ?", id).Delete(&Account{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}

// FindAdminUser 后台登入
func FindAdminUser(user Account, info LoginInfo) response.Response {
	var (
		result response.Response
	)
	if Db().Where("phone = ? AND password = ?", user.Name, user.Password).First(&user).RecordNotFound() {
		// 未找到记录
		return response.JSONError("账号或密码错误")
	}
	info.AccountID = user.ID
	InsertLoginInfo(info)
	token := utils.CreateToken(user.ID, user.Type)
	db.RedisInit().Set(user.ID, token, 60*24*15*60*time.Second)
	result = response.JSON(AdminJwtToken{
		Token: token,
		Type:  user.Type,
	})
	return result
}
