package service

import (
	"lottery-api/dao"
	"lottery-api/model"
)

type UserService struct{}

var userDao = new(dao.UserDao)

// SelectByLogin 登入
func (uService *UserService) SelectByLogin(user model.User) (*model.User, *model.SQLError) {
	res, err := userDao.SelectByLogin(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Select 查询
func (uService *UserService) Select(userID string) (*model.User, *model.SQLError) {
	return userDao.Select(userID)
}
