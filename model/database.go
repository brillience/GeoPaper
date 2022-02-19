package model

import (
	"geopaper/etc"
	"log"
	"runtime"
	"time"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	rdbPool *redis.Pool
)

func init() {
	// 连接mysql
	var err error
	db, err = gorm.Open(mysql.Open(etc.Config.Mysql.Datasource), &gorm.Config{})
	if err != nil {
		log.Fatalln("Connecting mysql error: ", err.Error())
	}
	// 创建表结构
	err = db.AutoMigrate(&FieldsOfStudy{}, &FieldOfStudyChildren{}, &GeoPaper{})
	if err != nil {
		log.Fatalln("Initting database table error: ", err.Error())
	}
	// 连接redis
	rdbPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				etc.Config.Redis.Addr,
				redis.DialDatabase(etc.Config.Redis.Db),
				redis.DialReadTimeout(time.Duration(100)*time.Second),
				redis.DialWriteTimeout(time.Duration(100)*time.Second),
				redis.DialConnectTimeout(time.Duration(100)*time.Second),
				redis.DialPassword(etc.Config.Redis.Password),
			)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     runtime.GOMAXPROCS(runtime.NumCPU()),
		MaxActive:   runtime.GOMAXPROCS(runtime.NumCPU()),
		IdleTimeout: time.Duration(100) * time.Second,
		Wait:        true,
	}
}
