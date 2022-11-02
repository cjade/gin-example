// Package models
// @Description: 书架
// @Author: Jade
// @Date: 2022/10/31 17:47
package models

type Bookshelf struct {
	BookshelfId uint `gorm:"primaryKey"`
	BookId      uint
	UserId      uint
	Model
}
