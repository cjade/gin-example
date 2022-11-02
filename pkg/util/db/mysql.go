package db

import (
	"fmt"
	"gin-example/init/config"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Mysql *gorm.DB

func init() {
	host := config.Cfg.Mysql.Host
	port := config.Cfg.Mysql.Port
	user := config.Cfg.Mysql.User
	pwd := config.Cfg.Mysql.Password
	tablePrefix := config.Cfg.Mysql.TablePrefix

	dataBase := config.Cfg.Mysql.DbName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, dataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})

	if config.Cfg.Debug {
		// 打印sql
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	if err != nil {
		log.Fatalln("Open mysql error : ", err.Error())
	}

	Mysql = db
}
