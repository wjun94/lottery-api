package dao

import (
	"job-api/model"
)

type MsgDao struct{}

// Insert 添加留言
func (dao *MsgDao) Insert(data model.Msg) *model.SQLError {
	if dbc := Db().Create(&data); dbc.Error != nil {
		dbc.Rollback()
		return GetError(dbc)
	}
	return nil
}

// SelectList 查询留言
func (dao *MsgDao) SelectList(current uint16, pageSize uint16, userID string) ([]model.Msg, uint32, *model.SQLError) {
	var info []model.Msg
	var count uint32
	if dbc := Db().Model(&model.Msg{}).Order("read_at desc").Where("user_id = ?", userID).Count(&count).Limit(pageSize).Offset((current - 1) * pageSize).Find(&info); dbc.Error != nil {
		return info, 0, GetError(dbc)
	}
	return info, count, nil
}
