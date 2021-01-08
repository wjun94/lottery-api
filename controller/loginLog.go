package controller

import (
	"job-api/response"
	"job-api/service"

	"github.com/gin-gonic/gin"
)

var loginLogService = new(service.LoginLogService)

// ListLoginLog 登入日志
func ListLoginLog(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	current, pageSize := GetPageParams(c)
	res, count, err := loginLogService.SelectList(current, pageSize, info.UserID)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultPageJSON(c, res, count, 1)
}
