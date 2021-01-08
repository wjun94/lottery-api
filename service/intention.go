package service

import (
	"job-api/dao"
	"job-api/model"
)

// IntentionService 服务
type IntentionService struct{}

var intentionDao = new(dao.IntentionDao)

// Insert 添加
func (service *IntentionService) Insert(data model.Intention) (*string, *model.SQLError) {
	return intentionDao.Insert(data)
}

// Select 查询
func (service *IntentionService) Select(id string) (*model.Intention, *model.SQLError) {
	return intentionDao.Select(id)
}

// Update 更新
func (service *IntentionService) Update(data model.UIntention) *model.SQLError {
	return intentionDao.Update(data)
}
