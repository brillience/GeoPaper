package model

import (
	"geopaper/etc"
	"log"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	// 连接mysql
	var err error
	db, err = gorm.Open(mysql.Open(etc.Config.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalln("Connecting mysql error: ", err.Error())
	}
	// 创建表结构
	err = db.AutoMigrate(&FieldsOfStudy{}, &FieldOfStudyChildren{}, &GeoPaper{})
	if err != nil {
		log.Fatalln("Initting database table error: ", err.Error())
	}
	// 连接redis
	redis.NewPool
}
