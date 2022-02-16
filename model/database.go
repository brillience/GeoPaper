package model

import (
	"geopaper/etc"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	rdb *redis.Client
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
	rdb = redis.NewClient(&redis.Options{
		Addr:     etc.Config.Redis.Addr,
		Password: etc.Config.Redis.Password,
		DB:       etc.Config.Redis.DB,
	})
}
