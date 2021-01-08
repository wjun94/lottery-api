package service

import (
	"job-api/dao"
	"job-api/model"
)

type UserService struct{}

var userDao = new(dao.UserDao)

// Select 查询
func (uService *UserService) Select(userID string) (*model.User, *model.SQLError) {
	return userDao.Select(userID)
}

// Insert 插入用户数据
func (uService *UserService) Insert(user model.User) (model.User, *model.SQLError) {
	return userDao.Insert(user)
}

// SelectByLogin 登入
func (uService *UserService) SelectByLogin(user model.User, loginLog model.LoginLog) (*model.User, *model.SQLError) {
	res, err := userDao.SelectByLogin(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdatePwd 更新密码Service
func (uService *UserService) UpdatePwd(userID string, pwd string, newPwd string) *model.SQLError {
	return userDao.UpdatePwd(userID, pwd, newPwd)
}
