package controller

import (
	"job-api/model"
	"job-api/response"
	"job-api/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var msgService = new(service.MsgService)

// CreateMsg 创建留言
func CreateMsg(c *gin.Context) {
	type verify struct {
		Name   string `json:"name" validate:"required" label:"标题"`
		Desc   string `json:"desc" validate:"required" label:"描述"`
		Friend string `json:"friend" form:"friend" label:"对方的ID"`
		ReadAt string `json:"readAt" label:"阅读时间"`
		Phone  string `json:"phone" validate:"required,numeric,min=6,max=11" label:"联系手机"`
		Status bool   `json:"status" label:"留言状态"`
	}
	var (
		vf  verify
		msg model.Msg
	)
	// c.ShouldBind(&vf)
	c.ShouldBindBodyWith(&vf, binding.JSON)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	c.ShouldBindBodyWith(&msg, binding.JSON)
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	msg.UserID = info.UserID
	msg.Status = false
	msg.SendAt = time.Now().Format("2006-01-02 15:04:05")
	msg.ID = Utils.UUID()
	err := msgService.Create(msg)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultSuccess(c)
}

// ListMsg 列表
func ListMsg(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	current, pageSize := GetPageParams(c)
	res, count, err := msgService.SelectList(current, pageSize, info.UserID)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultPageJSON(c, res, count, 1)
}
