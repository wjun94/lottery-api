package controller

import (
	"fmt"
	"lottery-api/model"
	"lottery-api/response"
	"lottery-api/service"

	"github.com/gin-gonic/gin"
)

var lotteryService = new(service.LotteryService)

// LotteryList 房价后台列表
func LotteryList(c *gin.Context) {
	current, pageSize := GetPageParams(c)
	var selectLotList model.SelectLotList
	c.ShouldBind(&selectLotList)
	fmt.Println(selectLotList)
	res, count, err := lotteryService.SelectList(selectLotList, current, pageSize)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	// var list []model.Lottery
	// rp, _ := json.Marshal(res)
	// json.Unmarshal(rp, &list)
	response.ResultPageJSON(c, res, count, current)
}
