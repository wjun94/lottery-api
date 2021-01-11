package app

/**
 * 描述信息：路由
 */
import (
	"lottery-api/config"
	"lottery-api/controller"
	"lottery-api/db"
	"lottery-api/service"
	"lottery-api/utils"

	"github.com/gin-gonic/gin"
)

var Utils = new(utils.Utils)

// InitApp 初始化
func InitApp() {
	println("version:0.0.3")
	r := gin.Default()
	v := r.Group(config.Config.Route.PathPrefix)
	v.POST("/login", controller.Login)

	v.Use(Middleware()) // 下面需要token认证
	v.DELETE("/loginout", controller.Loginout)
	{
		v.GET("/lotteryList", controller.LotteryList)
	}
	r.Run(":7001")
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
