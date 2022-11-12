// Package repositories
// @Description:
// @Author: Jade
// @Date: 2022/10/31
package repositories

import (
	"gin-example/internal/models"
	"gin-example/pkg/util/db"
)

type Book struct {
	models.Books
}

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
// @Return  uint64
// @Return  error
func (b *Book) Create() (uint, error) {
	result := db.Mysql.Create(&b)
	return b.BookId, result.Error
}
