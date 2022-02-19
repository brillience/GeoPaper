package etc

import (
	"log"

	"github.com/jinzhu/configor"
)

type Configor struct {
	Appname string `default:"GeoPaper"`
	Mysql   struct {
		Datasource string `default:"root:WQAOIaiona8X@tcp(127.0.0.1:33069)/geopaper?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"`
	}
	Redis struct {
		Addr     string `default:"localhost:63799"`
		Password string `default:"WQAOIaiona8X"`
		Db       int    `default:"0"`
	}
}

var Config Configor

func init() {
	err := configor.Load(&Config, "etc/config.yaml")
	if err != nil {
		log.Fatalln("Loading config.yaml error : ", err.Error())
	}
	log.Println("Loading config.yaml success!")
}
