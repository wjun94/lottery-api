package dao

import "job-api/model"

// IntentionDao 结构体
type IntentionDao struct{}

// Insert 插入数据
func (dao *IntentionDao) Insert(data model.Intention) (*string, *model.SQLError) {
	if dbc := Db().Create(&data); dbc.Error != nil {
		dbc.Rollback()
		return nil, GetError(dbc)
	}
	return &data.ID, nil
}

// Select 插入数据
func (dao *IntentionDao) Select(id string) (*model.Intention, *model.SQLError) {
	var info model.Intention
	if dbc := Db().Where("id = ?", id).Find(&info); dbc.Error != nil {
		dbc.Rollback()
		return nil, GetError(dbc)
	}
	return &info, nil
}

// Update 更新数据
func (dao *IntentionDao) Update(data model.UIntention) *model.SQLError {
	if dbc := Db().Table("intentions").Where("id = ?", data.ID).Updates(&data); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "更新失败!",
		}
	}
	return nil
}
