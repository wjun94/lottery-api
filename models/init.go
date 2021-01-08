package models

/**
 * 文档作者: wjun94
 * 创建时间：2019年10月17日
 * 修改时间：2019年10月18日
 * 描述信息：初始化，创建表和表外键
 */
import (
	"mango-api/config"
	"mango-api/db"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Dbs 数据库链接的参数
var Dbs *gorm.DB = db.MysqlConn()

// Db 使用数据库Use database_name
func Db() *gorm.DB {
	err := Dbs.Exec("USE " + config.Config.DB.Base)
	if err == nil {
		println("use 数据库报错了")
		Dbs = db.MysqlConn()
		Dbs.Exec("USE " + config.Config.DB.Base)
	}
	return Dbs
}

func create(models interface{}) {
	Db().AutoMigrate(models)
}

// 初始化，创建数据库和表
func init() {
	create(&User{})
	create(&Address{})
	create(&Commodity{})
	create(&Link{})
	if !Db().HasTable("accounts") {
		create(&Account{})
		Db().Create(&Account{ID: uuid.New().String(), Name: "admin", Password: "123456", Phone: "admin", Type: 2, CreateTime: "2020-01-01"})
	}
	create(&Coupon{})
	create(&UserCoupon{})
	create(&CommodityAd{})
	create(&Order{})
	create(&OrderCommodity{})
	create(&OrderAddress{})
	create(&Promote{})
	create(&Comment{})
	create(&LoginInfo{})
	Db().Model(&CommodityAd{}).AddForeignKey("commodity_id", "commodities(id)", "RESTRICT", "RESTRICT")
	Db().Model(&Address{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db().Model(&Link{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	Db().Model(&UserCoupon{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db().Model(&OrderCommodity{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	Db().Model(&OrderAddress{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	Db().Model(&Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db().Model(&LoginInfo{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
}
