package models

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月18日
 * 修改时间：2019年10月18日
 * 描述信息：用户信息数据库操作
 */
import (
	"mango-api/db"
	"mango-api/response"
	"mango-api/utils"
	"math/rand"
	"strings"
	"time"
)

// User 用户信息表数据结构
type User struct {
	ID         string `gorm:"primary_key" json:"id"`               // 用户唯一标识
	WxID       string `gorm:"unique" json:"wxId"`                  // 微信id
	Platform   byte   `gorm:"not null" json:"platform"`            // 哪个平台创建的账户(1: h5 2:小程序)[不要删]
	CreateTime string `gorm:"not null" json:"createTime"`          // 创建时间
	LinkID     string `gorm:"not null" json:"linkId"`              // 下单渠道ID
	AvatarURL  string `gorm:"not null" json:"avatarUrl,omitempty"` // 头像链接
	City       string `gorm:"not null" json:"city,omitempty"`      // 城市
	Country    string `gorm:"not null" json:"country,omitempty"`   // 国家
	Province   string `gorm:"not null" json:"province,omitempty"`  // 省
	Gender     byte   `gorm:"not null" json:"gender,omitempty"`    // 性别
	Language   string `gorm:"not null" json:"language,omitempty"`  // 语言
	NickName   string `gorm:"not null" json:"nickName,omitempty"`  // 用户名
	AddressID  string `gorm:"not null" json:"addressId,omitempty"` // 地址id
	Extend1    byte   `gorm:"not null" json:"extend1,omitempty"`   // 扩展字段1
	Extend2    string `gorm:"not null" json:"extend2,omitempty"`   // 扩展字段2
}

// JwtToken 用户token
type JwtToken struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}

// InsertUser 添加数据
func InsertUser(user User) response.Response {
	var (
		result response.Response
	)
	if Db().Where("wx_id = ?", user.WxID).First(&user).RecordNotFound() {
		// 未找到记录
		rand.Seed(time.Now().Unix())
		user.ID = utils.GetID()
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		if dbc := Db().Create(&user); dbc.Error != nil {
			// 返回错误信息
			if strings.Contains(dbc.Error.Error(), "1062") {
				InsertUser(user)
			} else {
				return response.JSONError(dbc.Error.Error())
			}
		}
	}
	token := utils.CreateWxToken(user.ID)
	result = response.JSON(map[string]string{
		"token": token,
		"id":    user.ID,
	})
	db.RedisInit().Set(user.ID, token, 60*24*15*60*time.Second)
	return result
}

// SelectUser 查询后台用户数据
func SelectUser(wxID string) response.Response {
	var (
		result response.Response
		user   User
	)
	if Db().Where("wx_id = ?", wxID).First(&user).RecordNotFound() {
		// 未找到记录
		return response.JSONError("没有查到用户信息")
	}
	token := utils.CreateWxToken(user.ID)
	result = response.JSON(map[string]string{
		"token": token,
	})
	db.RedisInit().Set(user.ID, token, 60*24*15*60*time.Second)
	return result
}
