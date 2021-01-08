package dao

import (
	"job-api/model"
)

// CInfoDao 结构体
type CInfoDao struct{}

// Select 查询数据
func (dao *CInfoDao) Select(userID string) (*model.CInfo, *model.SQLError) {
	var info model.CInfo
	if err := Db().Where("user_id = ?", userID).Find(&info); err.Error != nil {
		return nil, GetError(err)
	}
	return &info, nil
}

// Update 更新公司信息
func (dao *CInfoDao) Update(data model.VerUpdCoy) *model.SQLError {
	if dbc := Db().Model(&model.CInfo{}).Where("user_id = ?", data.UserID).Updates(data); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "更新失败",
		}
	}
	return nil
}
