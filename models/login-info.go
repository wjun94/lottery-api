package models

import (
	"mango-api/response"
	"github.com/google/uuid"
	"time"
)

// LoginInfo 后台登入信息
type LoginInfo struct {
	ID        string `gorm:"primary_key" json:"id"`
	AccountID string `json:"accountId"`                 // 账户id
	Address   string `json:"address"`                   // 地址
	LoginTime string `gorm:"not null" json:"loginTime"` // 登入时间(要加)
	IP        string `gorm:"not null" json:"ip"`        // 登入ip(要加)
}

// InsertLoginInfo 添加登入数据
func InsertLoginInfo(info LoginInfo) response.Response {
	var (
		result response.Response
	)
	info.ID = uuid.New().String()
	info.LoginTime = time.Now().Format("2006-01-02 15:04:05")
	if dbc := Db().Create(&info); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}
