package dao

import "lottery-api/model"

type LotteryDao struct{}

// SelectList 登入查询
func (dao *LotteryDao) SelectList(current uint16, pageSize uint16) ([]model.Lottery, uint32, *model.SQLError) {
	var lot []model.Lottery
	var count uint32
	if dbc := Db().Table("lotteries").Count(&count).Order("create_at desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&lot); dbc.Error != nil {
		return nil, count, GetError(dbc)
	}
	return lot, count, nil
}
