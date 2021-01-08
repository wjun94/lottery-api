package service

import (
	"job-api/dao"
	"job-api/model"
)

// UInfoService 用户信息Service
type UInfoService struct{}

var uInfoDao = new(dao.UInfoDao)

// Create 创建dao
func (uInfo *UInfoService) Create(info model.UInfo) (*string, *model.SQLError) {
	return uInfoDao.Insert(info)
}

// Update 更新
func (uInfo *UInfoService) Update(info model.UpdateUInfo) *model.SQLError {
	return uInfoDao.Update(info)
}

// Select 查询dao
func (uInfo *UInfoService) Select(userID string) (*model.UInfo, *model.SQLError) {
	return uInfoDao.Select(userID)
}
