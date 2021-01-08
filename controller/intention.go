package controller

import (
	"job-api/model"
	"job-api/response"
	"job-api/service"

	"github.com/gin-gonic/gin"
)

var intentionService = new(service.IntentionService)

// CreateIntention 创建
func CreateIntention(c *gin.Context) {
	var vf model.Intention
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
	vf.ID = info.UserID
	id, err := intentionService.Insert(vf)
	response.ResultErrJSON(c, map[string]string{"id": *id}, err)
}

// SelectIntention 创建
func SelectIntention(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	res, err := intentionService.Select(info.UserID)
	response.ResultErrJSON(c, res, err)
}

// UpdateIntention 更新
func UpdateIntention(c *gin.Context) {
	var vf model.UIntention
	c.ShouldBindJSON(&vf)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	err := intentionService.Update(vf)
	response.ResultErrSuccess(c, err)
}
