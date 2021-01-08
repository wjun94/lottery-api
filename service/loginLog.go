package service

import (
	"job-api/dao"
	"job-api/model"
)

// LoginLogService 结构体
type LoginLogService struct{}

var loginLogDao = new(dao.LoginLogDao)

// Create 添加登入信息
func (service *LoginLogService) Create(userInfo model.LoginLog) {
	loginLogDao.Insert(userInfo)
}

// DeleteOld 删除半年前的数据
func (service *LoginLogService) DeleteOld(userID string) {
	loginLogDao.DeleteOld(userID)
}

// SelectList 查询登入信息
func (service *LoginLogService) SelectList(current uint16, pageSize uint16, userID string) ([]model.LoginLog, uint32, *model.SQLError) {
	return loginLogDao.SelectList(current, pageSize, userID)
}
