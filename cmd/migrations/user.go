// Package commands
// @Description:
// @Author: Jade
// @Date: 2022/10/31 17:35
package main

import (
	"fmt"
	"gin-example/internal/models"
	"gin-example/internal/repositories"
	"gin-example/pkg/util/db"
	"gin-example/pkg/util/hash"
	"log"
)

func aa() {
	//Setup()
	//return
	// 条件更新
	db.Mysql.Model(&models.Users{}).Where("id = ?", 1587723805004926976).Update("user_name", "aba")
	return
	u := &repositories.Users{}
	//u := models.Users{UserName: "ds"}
	u.UserName = "dsad"
	u.Password = hash.BcryptHash("123456")
	create, err := u.Create()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(create)

}

func Setup() {
	users := &models.Users{}

	// 删除表
	//s := db.Mysql.Migrator().DropTable(users).Error()
	//fmt.Println("删除表 ：", s)

	// 判断表是否存在
	if !db.Mysql.Migrator().HasTable(users) {
		// 不存在就创建
		s := db.Mysql.Debug().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表信息'").Migrator().CreateTable(users).Error()
		fmt.Println("创建表 ：", s)
	}

}
