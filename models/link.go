package models

import (
	"mango-api/response"
)

// Link h5链接
type Link struct {
	ID        string `gorm:"primary_key" json:"id,omitempty"`
	AccountID string `gorm:"not null" json:"accountId"`
	Platform  byte   `gorm:"not null" json:"platform"` // 哪个平台 1 h5 2 微信
	Name      string `gorm:"not null" json:"name"`     // 平台名
	Path      string `gorm:"not null" json:"path"`     // 平台链接
	Ad        string `json:"ad"`                       // 广告位
	Other     string `json:"other"`                    // 首页位
	QrCode    string `json:"qrCode"`                   // 微信二维码
}

// InsertLink 添加链接
func InsertLink(link Link) response.Response {
	var (
		result response.Response
	)
	// if link.Platform == 2 {
	// 	// 微信小程序生成平台连接
	// 	fmt.Println(GetQRCode(link.ID))
	// }
	if dbc := Db().Create(&link); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSON(link)
	}
	return result
}

// FindAllLink 获取所有渠道链接
func FindAllLink(current int, pageSize int) response.Response {
	var links []Link
	Db().Limit(pageSize).Offset((current - 1) * pageSize).Find(&links)
	return response.JSON(links)
}

// SelectLink 获取渠道链接
func SelectLink(id string) response.Response {
	var links []Link
	Db().Where("account_id = ?", id).Find(&links)
	return response.JSON(links)
}

// DeleteLink 删除链接
func DeleteLink(id string) response.Response {
	if dbc := Db().Where("id = ?", id).Delete(&Link{}); dbc.Error != nil {
		// 返回错误信息
		return response.JSONError(dbc.Error.Error())
	}
	return response.JSONSuccess()
}

// UpdateLink 更新数据
func UpdateLink(link Link) response.Response {
	var result response.Response
	if dbc := Db().Model(&Link{}).Where("id = ?", link.ID).Update(&link); dbc.Error != nil {
		result = response.JSONError(dbc.Error.Error())
	} else {
		result = response.JSONSuccess()
	}
	return result
}
