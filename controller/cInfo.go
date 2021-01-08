package controller

import (
	"job-api/model"
	"job-api/response"
	"job-api/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var cInfoService = new(service.CInfoService)

// CompanyInfo 企业用户信息
func CompanyInfo(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	data, err := cInfoService.Select(info.UserID)
	if err != nil {
		response.ResultSQLError(c, 500, err.Message)
		return
	}
	response.ResultJSON(c, data)
}

// UpdateCompanyInfo 更新企业信息
func UpdateCompanyInfo(c *gin.Context) {
	var (
		vf model.VerUpdCoy
	)
	c.ShouldBindBodyWith(&vf, binding.JSON)
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
	cInfoService.Update(vf)
	response.ResultSuccess(c)
}
