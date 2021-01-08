package models

import (
	"mango-api/response"
)

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月25日
 * 修改时间：2019年10月25日
 * 描述信息：广告位数据库操作
 */

type CommodityAd struct {
	CommodityID string `json:"commodityID"` // 商品id
}

// FindCommodityAdById 查询
func FindCommodityAdById(id string) bool {
	var commodityAd CommodityAd
	// 关联查询
	if dbc := Db().Select("*").Where("commodity_id = ?", id).Find(&commodityAd); dbc.Error != nil {
		// 返回错误信息
	}
	// 判断uuid不为空
	return commodityAd.CommodityID != ""
}

// FindCommodityAd 查询
func FindCommodityAd() response.Response {
	var commodityAd CommodityAd
	// 关联查询
	if dbc := Db().First(&commodityAd); dbc.Error != nil {
		// 返回错误信息
	}
	var res = map[string]string{"commodityID": commodityAd.CommodityID}
	// 判断uuid不为空
	return response.JSON(res)
}
