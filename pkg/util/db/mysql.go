package db

import (
	"fmt"
	"gin-example/configs"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *gorm.DB

func init() {
	host := configs.Cfg.Mysql.Host
	port := configs.Cfg.Mysql.Port
	user := configs.Cfg.Mysql.User
	pwd := configs.Cfg.Mysql.Password
	tablePrefix := configs.Cfg.Mysql.TablePrefix

	dataBase := configs.Cfg.Mysql.DbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})

	if configs.Cfg.Debug {
		// 打印sql
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	if err != nil {
		log.Fatalln("Open mysql error : ", err.Error())
	}

	Mysql = db
}
