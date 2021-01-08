package dao

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：公司表(存放账号、密码)
 */
import (
	"job-api/model"
)

// UserDao dao结构体
type UserDao struct{}

// Select 查询详情数据
func (dao *UserDao) Select(userID string) (*model.User, *model.SQLError) {
	var us model.User
	if dbc := Db().Where("id = ?", userID).Find(&us); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	Db().Model(&us).Related(&us.CInfo).Related(&us.UInfos)
	return &us, nil
}

// Insert 添加数据
func (dao *UserDao) Insert(user model.User) (model.User, *model.SQLError) {
	if dbc := Db().Create(&user); dbc.Error != nil {
		return user, GetError(dbc)
	}
	return user, nil
}

// SelectByLogin 登入查询
func (dao *UserDao) SelectByLogin(user model.User) (*model.User, *model.SQLError) {
	var us model.User
	if dbc := Db().Where("pwd = ? AND phone = ?", user.Pwd, user.Phone).Find(&us); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &us, nil
}

// UpdatePwd 更新密码数据库操作
func (dao *UserDao) UpdatePwd(userID string, pwd string, newPwd string) *model.SQLError {
	if dbc := Db().Model(&model.User{}).Where("id = ? AND pwd = ? ", userID, pwd).Update("pwd", Utils.Encrypt(newPwd)); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "原密码错误",
		}
	}
	return nil
}
