package service

import (
	"job-api/dao"
	"job-api/model"
)

type MsgService struct{}

var msgDao = new(dao.MsgDao)

// Create 创建
func (msg *MsgService) Create(data model.Msg) *model.SQLError {
	return msgDao.Insert(data)
}

// SelectList 查询用户留言
func (msg *MsgService) SelectList(current uint16, pageSize uint16, userID string) ([]model.Msg, uint32, *model.SQLError) {
	return msgDao.SelectList(current, pageSize, userID)
}
