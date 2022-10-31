// Package Repositories
// @Description:
// @Author: Jade
// @Date: 2022/10/31
package Repositories

import (
	"gin-example/app/models"
	"gin-example/utils/db"
)

type Book models.Books

// Create
//
// @Description: 创建
//
// @Author: Jade
//
// @Date: 2022-10-31 16:31:48
//
// @Receiver  b
//
//
// @Return  uint64
// @Return  error
func (b *Book) Create() (uint64, error) {
	result := db.Mysql.Create(&b)
	return b.BookId, result.Error
}
