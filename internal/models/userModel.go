package models

import (
	"gin-example/pkg/sf"
	"gorm.io/gorm"
)

type Users struct {
	UserId      uint64 `gorm:"primaryKey;comment:用户雪花ID"`
	UserName    string `gorm:"not null;size:50;comment:用户名称"`
	Password    string `gorm:"not null;type:varchar(60);comment:密码"`
	RealName    string `gorm:"not null;default:'';size:50;comment:真实名称"`
	Gender      int8   `gorm:"not null;default:0;comment:性别"`
	Email       string `gorm:"not null;default:'';type:varchar(100);comment:邮箱"`
	PhoneNumber string `gorm:"not null;default:'';size:11;comment:手机号码"`
	Age         uint8  `gorm:"not null;default:0;comment:年龄"`
	UserStatus  int8   `gorm:"not null;default:0;comment:用户状态"`
	Model
	Bookshelf []Bookshelf `gorm:"foreignKey:user_id"`
}

// BeforeCreate
//
// @Description: user_id 使用雪花ID
//
// @Author: Jade
//
// @Date: 2022-10-31 22:23:09
//
// @Receiver  u
//
// @Param  tx
//
// @Return  err
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := sf.GenId()
	u.UserId = uint64(id)
	return
}

// DeletedAtDefault
//
// @Description: 没有删除的用户
//
// @Author: Jade
//
// @Date: 2022-11-01 15:44:52
//
// @Param  db
//
// @Return  *gorm.DB
func DeletedAtDefault(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at = ?", "1970-01-01 08:00:00")
}

// DeletedAtNotDefault
//
// @Description: 删除的用户
//
// @Author: Jade
//
// @Date: 2022-11-01 15:53:17
//
// @Param  db
//
// @Return  *gorm.DB
func DeletedAtNotDefault(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at <> ?", "1970-01-01 08:00:00")
}
