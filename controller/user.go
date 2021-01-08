package controller

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月12日
 * 修改时间：2019年10月12日
 * 描述信息：登入
 */
import (
	"net/http"
	"time"

	"job-api/db"
	"job-api/model"
	"job-api/response"
	"job-api/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var userService = new(service.UserService)
var lService = new(service.LoginLogService)

// RegistComp 招聘者注册(企业)
func RegistComp(c *gin.Context) {
	type verify struct {
		Name  string `form:"name" validate:"required" label:"公司名"`
		Ind   string `form:"ind" validate:"required" label:"所属行业"`
		Prov  string `form:"prov" validate:"required" label:"省"`
		City  string `form:"city" validate:"required" label:"城市"`
		Addr  string `form:"addr" validate:"required" label:"详细地址"`
		Scale byte   `form:"scale" validate:"required" label:"企业规模"`
		Cont  string `form:"cont" validate:"required" label:"联系人"`
		Phone string `form:"phone" validate:"required" label:"手机号"`
		Email string `form:"email" validate:"required,email" label:"邮箱"`
	}
	var (
		user model.User
		vf   verify
	)
	c.ShouldBindBodyWith(&vf, binding.JSON)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, 400, *err)
		return
	}
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.ShouldBindBodyWith(&user.CInfo, binding.JSON)
	user.ID = Utils.UUID()
	user.Level = 3
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	user.Pwd = Utils.Encrypt("123456")
	user.CInfo.ID = Utils.UUID()
	user.CInfo.UserID = user.ID
	user.CInfo.Level = "t1"
	_, err := userService.Insert(user)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	response.ResultSuccess(c)
}

// Regist 求职者注册
func Regist(c *gin.Context) {
	var user model.User
	var verify model.VerifyUserRegist
	user.ID = Utils.UUID()
	user.Level = 4
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.ShouldBindBodyWith(&verify, binding.JSON)
	user.Pwd = Utils.Encrypt(user.Pwd)
	if err := Utils.Trans(verify); err != nil {
		response.ResultSQLError(c, 500, *err)
		return
	}
	_, err := userService.Insert(user)
	if err != nil {
		response.ResultSQLError(c, 500, err.Message)
		return
	}
	response.ResultSuccess(c)
}

// LoginCompany 企业后台登入
func LoginCompany(c *gin.Context) {
	type verify struct {
		Name string `form:"name" validate:"required" label:"账号"`
		Pwd  string `form:"pwd" validate:"required" label:"密码"`
	}
	var (
		user     model.User
		loginLog model.LoginLog
		vf       verify
	)
	c.ShouldBindBodyWith(&vf, binding.JSON)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, http.StatusBadRequest, *err)
		return
	}
	c.ShouldBindBodyWith(&user, binding.JSON)
	user.Phone = vf.Name
	user.Pwd = Utils.Encrypt(user.Pwd)
	loginLog.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	loginLog.IP = c.ClientIP()
	loginLog.ID = Utils.UUID()
	loginLog.Platform = Utils.GetPlatform(c)
	loginLog.Browser = Utils.GetBrowser(c)
	us, err := userService.SelectByLogin(user, loginLog)
	if err != nil || us.Level != 3 {
		response.ResultSQLError(c, 505, "")
		return
	}
	lService.Create(loginLog)
	lService.DeleteOld(us.ID)
	token := Utils.Create(us.ID, us.Level)
	db.RedisInit().Set(us.ID, token, 60*24*15*60*time.Second)
	response.ResultJSON(c, token)
}

// Login 求职者网登入
func Login(c *gin.Context) {
	var user model.User
	var verify model.VerifyUserLogin
	var loginLog model.LoginLog
	c.ShouldBindBodyWith(&verify, binding.JSON)
	if err := Utils.Trans(verify); err != nil {
		response.ResultSQLError(c, 500, *err)
		return
	}
	c.ShouldBindBodyWith(&user, binding.JSON)
	user.Pwd = Utils.Encrypt(user.Pwd)
	loginLog.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	loginLog.IP = c.ClientIP()
	loginLog.ID = Utils.UUID()
	loginLog.Platform = Utils.GetPlatform(c)
	loginLog.Browser = Utils.GetBrowser(c)
	us, err := userService.SelectByLogin(user, loginLog)
	loginLog.UserID = us.ID
	if err != nil || us.Level != 4 {
		response.ResultSQLError(c, 505, "账号或密码错误")
		return
	}
	lService.Create(loginLog)
	lService.DeleteOld(us.ID)
	token := Utils.Create(us.ID, us.Level)
	db.RedisInit().Set(us.ID, token, 60*24*15*60*time.Second)
	response.ResultJSON(c, token)
}

// Loginout 退出登录
func Loginout(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	db.RedisInit().Del(info.UserID)
	response.ResultSuccess(c)
}

// Detail 求职者或者企业信息信息
func Detail(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	res, err := userService.Select(info.UserID)
	if err != nil {
		response.ResultSQLError(c, err.Number, err.Message)
		return
	}
	switch res.Level {
	case 3:
		response.ResultJSON(c, res.UInfos)
	case 4:
		response.ResultJSON(c, res.CInfo)
	}
}

// UpdatePwd 更新密码
func UpdatePwd(c *gin.Context) {
	type verify struct {
		Pwd    string `form:"pwd" binding:"required" validate:"required" label:"原密码"`
		NewPwd string `form:"newPwd" validate:"required,nefield=Pwd,min=6,max=20" label:"新密码"`
		VerPwd string `form:"verPwd" validate:"required,eqfield=NewPwd" label:"确认新密码"`
	}
	var (
		vf verify
	)
	c.ShouldBind(&vf)
	if err := Utils.Trans(vf); err != nil {
		response.ResultSQLError(c, http.StatusBadRequest, *err)
		return
	}
	userInfo, _ := c.Get("userInfo")
	info := userInfo.(struct {
		UserID string
		Level  byte
	})
	if err := userService.UpdatePwd(info.UserID, Utils.Encrypt(vf.Pwd), vf.NewPwd); err != nil {
		response.ResultSQLError(c, 500, err.Message)
		return
	}
	response.ResultSuccess(c)
}
