package dao

/**
 * 描述信息：公司登入信息表
 */
import (
	"job-api/model"
	"time"
)

// LoginLogDao 结构体
type LoginLogDao struct{}

// Insert 插入用户登入数据
func (dao *LoginLogDao) Insert(data model.LoginLog) string {
	if dbc := Db().Create(&data); dbc.Error != nil {
		dbc.Rollback()
		return dbc.Error.Error()
	}
	return ""
}

// DeleteOld 删除半年前数据
func (dao *LoginLogDao) DeleteOld(userID string) {
	nowTime := time.Now().AddDate(0, 0, -180)
	Db().Where("create_at < ?", nowTime).Delete(&model.LoginLog{})
}

// SelectList 登入信息列表数据
func (dao *LoginLogDao) SelectList(current uint16, pageSize uint16, userID string) ([]model.LoginLog, uint32, *model.SQLError) {
	var info []model.LoginLog
	var count uint32
	if dbc := Db().Model(&model.LoginLog{}).Order("create_at desc").Where("user_id = ?", userID).Count(&count).Limit(pageSize).Offset((current - 1) * pageSize).Find(&info); dbc.Error != nil {
		return info, 0, GetError(dbc)
	}
	return info, count, nil
}
