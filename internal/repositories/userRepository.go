// Package repositories
// @Description:
// @Author: Jade
// @Date: 2022/10/31
package repositories

import (
	"errors"
	"gin-example/internal/models"
	"gin-example/pkg/util/db"
	"gorm.io/gorm"
	"log"
)

type UsersRepository struct {
	BaseRepository
}

var Users = models.Users{}

func (ur *UsersRepository) SetModel() *UsersRepository {
	ur.Model = &Users
	return ur
}

// Create
//
// @Description:  创建
//
// @Author: Jade
//
// @Date: 2022-10-31 16:02:32
//
// @Receiver  u
//
// @Return  int64
// @Return  error
func (ur *UsersRepository) Create(u models.Users) (uint64, error) {
	result := db.Mysql.Create(&u)
	return u.UserId, result.Error
}

// GetUserByAccount
//
// @Description: 通过账号获取用户
//
// @Author: Jade
//
// @Date: 2022-11-03 17:25:01
//
// @Receiver  u
//
// @Return  error
func (ur *UsersRepository) GetUserByAccount(userName string) models.Users {
	user := models.Users{}
	err := db.Mysql.Select(ur.Fields).
		Where("user_name = ?", userName).
		Or("email = ?", userName).
		Or("phone_number", userName).
		First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
	log.Printf("%v", user)
	return user
}
