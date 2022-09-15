package dao

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UserWriter IUserWriter
var UserReader IUserReader

func InitDataSources() {

	dsn_r := "douyin_demo_read:ATaleOf2C@tcp(127.0.0.1:3306)/douyin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db_r, err := gorm.Open(mysql.Open(dsn_r), &gorm.Config{})
	if err != nil {
		log.Infof("%v", err)
		panic(err)
	} else {
		UserReader = &mysqlUserReader{db_r}
		log.Infoln("get connected to douyin_demo_read")

	}

	dsn_w := "douyin_demo_write:ATaleOf2C@tcp(127.0.0.1:3306)/douyin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db_w, err := gorm.Open(mysql.Open(dsn_w), &gorm.Config{})
	if err != nil {
		log.Infof("%v", err)
		panic(err)
	} else {
		UserWriter = &mysqlUserWriter{db_w}
		log.Infoln("get connected to douyin_demo_write")

	}

}
