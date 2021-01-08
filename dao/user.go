package dao

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：公司表(存放账号、密码)
 */
import (
	"lottery-api/model"
)

// UserDao dao结构体
type UserDao struct{}

// SelectByLogin 登入查询
func (dao *UserDao) SelectByLogin(user model.User) (*model.User, *model.SQLError) {
	var us model.User
	if dbc := Db().Where("pwd = ? AND phone = ?", user.Pwd, user.Phone).Find(&us); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &us, nil
}

// Select 查询详情数据
func (dao *UserDao) Select(userID string) (*model.User, *model.SQLError) {
	var us model.User
	if dbc := Db().Where("id = ?", userID).Find(&us); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &us, nil
}
