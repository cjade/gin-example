// Package models
// @Description:
// @Author: Jade
// @Date: 2022/10/31 21:37
package models

import (
	"time"
)

type Model struct {
	//ID        uint      `gorm:"primaryKey"`
	DeletedAt time.Time `gorm:"type:datetime(0);index;not null;default:'1970-01-01 08:00:00';comment:删除时间"`
	CreatedAt time.Time `gorm:"type:datetime(0);NOT NULL;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime(0);NOT NULL;comment:更新时间"`
}
