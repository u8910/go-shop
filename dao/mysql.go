package dao

import (
	"go_shop/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var config = conf.Conf()

func MysqlInit() {

	dsn := config.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库启动失败")
	}
	DB = db
}
