package dao

import (
	"lottery-api/config"
	"lottery-api/db"
	"lottery-api/model"
	"lottery-api/utils"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Utils = new(utils.Utils)

// Dbs 数据库链接的参数
var Dbs *gorm.DB = db.MysqlInit()

// Db 使用数据库Use database_name
func Db() *gorm.DB {
	Dbs.Exec("USE " + config.Config.DB.Base)
	err := Dbs.Exec("USE " + config.Config.DB.Base)
	if err.Error != nil {
		println("use 数据库报错了")
		Dbs = db.MysqlInit()
		Dbs.Exec("USE " + config.Config.DB.Base)
	}
	return Dbs
}

// GetError 数据库错误处理
func GetError(dbc *gorm.DB) *model.SQLError {
	err := dbc.Error
	if gorm.IsRecordNotFoundError(err) {
		return &model.SQLError{Number: 200, Message: "空数据"}
	}
	b := Utils.HasOwnProperty(err, "Number")
	if b == true {
		mError := err.(*mysql.MySQLError)
		return &model.SQLError{Number: mError.Number, Message: mError.Message}
	} else {
		return Parse(err.Error())
	}
}

// Parse 解析String类型报错
func Parse(str string) *model.SQLError {
	if str == "record not found" {
		return &model.SQLError{Number: 204, Message: "没有查询到数据"}
	}
	arr := strings.Split(str, ": ")
	chiReg := regexp.MustCompile("[0-9]+")
	Number := chiReg.FindAllString(arr[0], -1)
	return &model.SQLError{Number: Utils.StrToUin16(Number[0]), Message: arr[1]}
}
