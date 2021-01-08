package app

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：路由
 */
import (
	"job-api/config"
	"job-api/controller"
	"job-api/db"
	"job-api/service"
	"job-api/utils"

	"github.com/gin-gonic/gin"
)

var Utils = new(utils.Utils)

// InitApp 初始化
func InitApp() {
	println("version:0.0.3")
	r := gin.Default()
	v := r.Group(config.Config.Route.PathPrefix)
	v.POST("/login", controller.Login)
	v.POST("/regist", controller.Regist)
	v.POST("/registCompany", controller.RegistComp)
	v.POST("/loginCompany", controller.LoginCompany)
	{
		v.GET("/listRcrtByComp", controller.ListRcrtByComp)
		v.GET("/rcrtDetail", controller.RcrtDetail)
	}
	v.Use(Middleware()) // 下面需要token认证
	v.DELETE("/loginout", controller.Loginout)
	// 求职者官网
	{
		v.POST("/createUInfo", controller.CreateUInfo)
		v.PATCH("/updateUInfo", controller.UpdateUInfo)
		v.GET("/getUInfo", controller.SelectUInfo)
	}
	{
		v.POST("/createIntention", controller.CreateIntention)
		v.GET("/getIntention", controller.SelectIntention)
		v.PATCH("/updateIntention", controller.UpdateIntention)
	}
	{
		v.POST("/createWork", controller.CreateWork)
	}
	// 企业网站官网
	{
		v.POST("/createRecruit", controller.CreateRecruit)
		v.GET("/listRecruit", controller.ListRecruit)
		v.GET("/getRecruit", controller.SelectRecruit)
		v.PATCH("/updateRecruit", controller.UpdateRecruit)
		v.PATCH("/updateRecruits", controller.UpdateRecruits)
		v.DELETE("/deleteRecruit", controller.DeleteRecruit)
	}
	{
		v.POST("/createMsg", controller.CreateMsg)
		v.GET("/listMsg", controller.ListMsg)
	}
	{
		v.PATCH("/updatePwd", controller.UpdatePwd)
		v.GET("/companyInfo", controller.CompanyInfo)
		v.PATCH("/updateCompanyInfo", controller.UpdateCompanyInfo)
	}
	{
		v.GET("/listLoginLog", controller.ListLoginLog)
	}
	{
		v.GET("/detail", controller.Detail)
	}
	r.Run()
}

// Middleware 中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		loginUserID, level := Utils.GetUserID(tokenStr)
		rd, _ := db.RedisInit().Get(loginUserID).Result()
		var uService = new(service.UserService)
		_, err := uService.Select(loginUserID)
		if tokenStr == "" || loginUserID == "" || err != nil {
			c.JSON(401, gin.H{"message": "请重新登入"})
			c.Abort()
		} else if rd != tokenStr {
			c.Abort()
			c.JSON(401, gin.H{"message": "请重新登入"})
		} else {
			// context.WithValue(c, "info", map[string]string{"UserID": loginUserID})
			c.Set("userInfo", struct {
				UserID string
				Level  byte
			}{loginUserID, level})
			c.Next()
		}
	}
}
