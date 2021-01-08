package controller

import (
	"job-api/model"
	"job-api/response"
	"job-api/service"

	"github.com/gin-gonic/gin"
)

var uInfoService = new(service.UInfoService)

// CreateUInfo 创建用户信息
func CreateUInfo(c *gin.Context) {
	var (
		vf model.UInfo
	)
	c.ShouldBindJSON(&vf)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})

	vf.ID = Utils.UUID()
	vf.UserID = info.UserID
	id, err := uInfoService.Create(vf)
	response.ResultErrJSON(c, map[string]string{"id": *id}, err)
}

// UpdateUInfo 更新用户信息
func UpdateUInfo(c *gin.Context) {
	var (
		vf model.UpdateUInfo
	)
	c.ShouldBindJSON(&vf)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})

	vf.UserID = info.UserID
	err := uInfoService.Update(vf)
	response.ResultErrSuccess(c, err)
}

// SelectUInfo 获取用户信息
func SelectUInfo(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	res, err := uInfoService.Select(info.UserID)
	response.ResultErrJSON(c, res, err)
}
