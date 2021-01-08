package controller

import (
	"lottery-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Utils 工具类
var Utils = new(utils.Utils)

// GetPageParams 获取页码
func GetPageParams(c *gin.Context) (uint16, uint16) {
	current := c.DefaultQuery("current", "1")    // 获取值并设置初始值
	pageSize := c.DefaultQuery("pageSize", "10") // 获取值并设置初始值
	x, _ := strconv.Atoi(current)
	y, _ := strconv.Atoi(pageSize)
	return uint16(x), uint16(y)
}
