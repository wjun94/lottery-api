package dao

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：公司招聘信息表
 */
import (
	"fmt"
	"job-api/model"
)

type RecruitDao struct{}

// Insert 招聘表添加数据
func (rcrt *RecruitDao) Insert(body model.Recruit) (model.Recruit, *model.SQLError) {
	if dbc := Db().Create(&body); dbc.Error != nil {
		dbc.Rollback()
		return body, GetError(dbc)
	}
	return body, nil
}

// SelectList 招聘信息列表数据
func (rcrt *RecruitDao) SelectList(current uint16, pageSize uint16, userID string) ([]model.Recruit, uint32, *model.SQLError) {
	var rt []model.Recruit
	var count uint32
	if dbc := Db().Model(&model.Recruit{}).Order("update_at desc").Where("user_id = ?", userID).Count(&count).Limit(pageSize).Offset((current - 1) * pageSize).Find(&rt); dbc.Error != nil {
		return rt, 0, GetError(dbc)
	}
	return rt, count, nil
}

// SelectListRcrtByComp 查询关联列表
func (rcrt *RecruitDao) SelectListRcrtByComp(current uint16, pageSize uint16) ([]model.ListRcrtByComp, uint32, *model.SQLError) {
	var rt []model.ListRcrtByComp
	var count uint32
	if dbc := Db().Table("recruits").Order("update_at desc").Preload("User").Preload("User.CInfo").Count(&count).Limit(pageSize).Offset((current - 1) * pageSize).Find(&rt); dbc.Error != nil {
		return rt, 0, GetError(dbc)
	}
	return rt, count, nil
}

// SelectDetail 查询关联列表
func (rcrt *RecruitDao) SelectDetail(id string) (model.RcrtDetail, *model.SQLError) {
	var rt model.RcrtDetail
	if dbc := Db().Table("recruits").Where("id = ?", id).Preload("User.CInfo").Find(&rt); dbc.Error != nil {
		return rt, GetError(dbc)
	}
	return rt, nil
}

// Select 查询招聘信息
func (rcrt *RecruitDao) Select(id string, userID string) (*model.Recruit, *model.SQLError) {
	var recruit model.Recruit
	if dbc := Db().Where("id = ? AND user_id = ?", id, userID).Find(&recruit); dbc.Error != nil {
		return nil, GetError(dbc)
	}
	return &recruit, nil
}

// Update 更新招聘信息
func (rcrt *RecruitDao) Update(recruit model.UpdateRcrt) *model.SQLError {
	if dbc := Db().Model(&model.Recruit{}).Where("id = ?", recruit.ID).Updates(&recruit); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "更新失败!",
		}
	}
	return nil
}

// Updates 批量更新
func (rcrt *RecruitDao) Updates(recruit map[string]interface{}) *model.SQLError {
	if dbc := Db().Model(&model.Recruit{}).Where("id IN (?)", recruit["ids"]).Updates(recruit); dbc.RowsAffected == 0 {
		return &model.SQLError{
			Message: "更新失败!",
		}
	}
	return nil
}

// Delete 删除职位
func (rcrt *RecruitDao) Delete(ids []string) *model.SQLError {
	fmt.Println(ids)
	if dbc := Db().Where("id IN (?)", ids).Delete(&model.Recruit{}); dbc.Error != nil {
		return GetError(dbc)
	}
	return nil
}
