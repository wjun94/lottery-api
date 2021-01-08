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

	"mango-api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// MysqlConn 连接数据库
func MysqlConn() *gorm.DB {
	// var baseInfo = fmt.Sprintf("%s:%s@tcp(localhost:3306)/?charset=utf8mb4,utf8&parseTime=true", config.Config.DB.User, config.Config.DB.Password)
	var baseInfo = fmt.Sprintf("%s:%s@/?charset=utf8mb4,utf8&parseTime=true", config.Config.DB.User, config.Config.DB.Password)
	db, err := gorm.Open("mysql", baseInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second)
	UseDatabase(db)
	// defer Db.Close()
	return db
}

// UseDatabase 创建数据库并使用
func UseDatabase(db *gorm.DB) *gorm.DB {
	db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Config.DB.Base)
	db.Exec("USE " + config.Config.DB.Base)
	return db
}
