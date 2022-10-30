package models

import (
	"gin-example/utils/db"
	"time"
)

type Books struct {
	BookId       uint64 `json:"book_id" gorm:"primaryKey"`
	TypeId       uint   `json:"type_id"`
	BookName     string `json:"book_name" gorm:"size:50"`
	Auth         string `json:"auth" gorm:"size:50"`
	Intro        string `json:"intro"`
	CoverPicture string `json:"cover_picture"`
	DeletedAt    uint8
	CreatedTime  time.Time `json:"created_time" gorm:"autoCreateTime" `
	UpdatedTime  time.Time `gorm:"autoUpdateTime"`
}

// Create
// @Description:
// @receiver b
// @return uint64  id
// @return error
func (b *Books) Create() (uint64, error) {
	result := db.Mysql.Create(&b)
	return b.BookId, result.Error
}
