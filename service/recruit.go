package service

import (
	"job-api/dao"
	"job-api/model"
)

type RecruitService struct{}

var recruitDao = new(dao.RecruitDao)

// Insert 创建招聘服务
func (rcrt *RecruitService) Insert(recruit model.Recruit) (model.Recruit, *model.SQLError) {
	return recruitDao.Insert(recruit)
}

// SelectListRcrtByComp 查询关联列表
func (rcrt *RecruitService) SelectListRcrtByComp(current uint16, pageSize uint16) ([]model.ListRcrtByComp, uint32, *model.SQLError) {
	return recruitDao.SelectListRcrtByComp(current, pageSize)
}

// SelectDetail 查询关联列表
func (rcrt *RecruitService) SelectDetail(id string) (model.RcrtDetail, *model.SQLError) {
	return recruitDao.SelectDetail(id)
}

// SelectList 查询列表数据
func (rcrt *RecruitService) SelectList(current uint16, pageSize uint16, userID string) ([]model.Recruit, uint32, *model.SQLError) {
	return recruitDao.SelectList(current, pageSize, userID)
}

// Select 查询
func (rcrt *RecruitService) Select(id string, userId string) (*model.Recruit, *model.SQLError) {
	return recruitDao.Select(id, userId)
}

// Update 更新
func (rcrt *RecruitService) Update(recruit model.UpdateRcrt) *model.SQLError {
	return recruitDao.Update(recruit)
}

// Updates 批量更新
func (rcrt *RecruitService) Updates(recruit map[string]interface{}) *model.SQLError {
	return recruitDao.Updates(recruit)
}

// Delete 删除职位
func (rcrt *RecruitService) Delete(ids []string) *model.SQLError {
	return recruitDao.Delete(ids)
}
