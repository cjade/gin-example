// Package repositories
// @Description:
// @Author: Jade
// @Date: 2022/10/31
package repositories

import (
	"gin-example/pkg/util/db"
	"gorm.io/gorm"
)

type BaseRepository struct {
	Model  interface{}
	Fields []string
}

func (br *BaseRepository) Find() *gorm.DB {
	return db.Mysql.Select(br.Fields).Find(&br.Model)
}
