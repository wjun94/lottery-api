package service

import (
	"lottery-api/dao"
	"lottery-api/model"
)

type LotteryService struct{}

var lotteryDao = new(dao.LotteryDao)

// SelectList 查询列表
func (service *LotteryService) SelectList(params model.SelectLotList, current uint16, pageSize uint16) ([]model.Lottery, uint32, *model.SQLError) {
	return lotteryDao.SelectList(params, current, pageSize)
}
