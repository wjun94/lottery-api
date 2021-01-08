package service

import (
	"job-api/dao"
	"job-api/model"
)

// CInfoService 企业用户service
type CInfoService struct{}

var cInfoDao = new(dao.CInfoDao)

// Select 查询
func (service *CInfoService) Select(userID string) (*model.CInfo, *model.SQLError) {
	return cInfoDao.Select(userID)
}

func (service *CInfoService) Update(data model.VerUpdCoy) *model.SQLError {
	return cInfoDao.Update(data)
}
