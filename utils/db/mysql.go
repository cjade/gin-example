package db

import (
	"fmt"
	"gin-example/conf"
	"gorm.io/gorm/schema"

	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *gorm.DB

func init() {
	host := conf.Cfg.Mysql.Host
	port := conf.Cfg.Mysql.Port
	user := conf.Cfg.Mysql.User
	pwd := conf.Cfg.Mysql.Password
	tablePrefix := conf.Cfg.Mysql.TablePrefix

	dataBase := conf.Cfg.Mysql.DbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalln("Open mysql error : ", err.Error())
	}

	Mysql = db
}
