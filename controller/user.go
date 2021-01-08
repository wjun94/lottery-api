package controller

import (
	"lottery-api/db"
	"lottery-api/model"
	"lottery-api/response"
	"lottery-api/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var userService = new(service.UserService)

// Login 登入
func Login(c *gin.Context) {
	var user model.User
	var verify model.VerifyLogin
	c.ShouldBindBodyWith(&verify, binding.JSON)
	if err := Utils.Trans(verify); err != nil {
		response.ResultSQLError(c, 500, *err)
		return
	}
	c.ShouldBindBodyWith(&user, binding.JSON)
	us, err := userService.SelectByLogin(user)
	if err != nil {
		response.ResultSQLError(c, 505, "账号或密码错误")
		return
	}
	token := Utils.CreateToken(us.ID, us.Level)
	db.RedisInit().Set(us.ID, token, 60*24*15*60*time.Second)
	response.ResultJSON(c, token)
}
