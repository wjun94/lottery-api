package db

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：链接数据库
 */
import (
	"fmt"
	"log"

	"lottery-api/config"
	"lottery-api/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 数据库
var DB *gorm.DB

// MysqlInit 连接数据库
func MysqlInit() *gorm.DB {
	var baseInfo = fmt.Sprintf("%s:%s@/?charset=utf8mb4,utf8&parseTime=true", config.Config.DB.User, config.Config.DB.Password)
	db, err := gorm.Open("mysql", baseInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second)
	UseDatabase(db)
	DB = db
	CreateTables()
	// defer Db.Close()
	return db
}

// UseDatabase 创建数据库并使用
func UseDatabase(db *gorm.DB) *gorm.DB {
	db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Config.DB.Base)
	db.Exec("USE " + config.Config.DB.Base)
	return db
}

func create(models interface{}) {
	DB.AutoMigrate(models)
	if !DB.HasTable(models) {
		DB.CreateTable(models)

	}
}

// CreateTables 初始化表
func CreateTables() {
	// user := model.User{ID: "1", Level: 1, Phone: "13588227124", Pwd: "123456"}
	// DB.Create(&user)
	create(&model.User{})
}
