package dao

import "job-api/model"

// UInfoDao 用户信息数据库操作
type UInfoDao struct{}

// Insert 添加数据
func (dao *UInfoDao) Insert(info model.UInfo) (*string, *model.SQLError) {
	if dbc := Db().Create(&info); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &info.ID, nil
}

// Update 更新数据
func (dao *UInfoDao) Update(info model.UpdateUInfo) *model.SQLError {
	if dbc := Db().Table("uinfos").Where("id = ?", info.ID).Updates(&info); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "更新失败!",
		}
	}
	return nil
}

// Select 查询数据
func (dao *UInfoDao) Select(userID string) (*model.UInfo, *model.SQLError) {
	var info model.UInfo
	if dbc := Db().Where("user_id = ?", userID).Find(&info); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &info, nil
}
