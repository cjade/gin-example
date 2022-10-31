// Package Repositories
// @Description:
// @Author: Jade
// @Date: 2022/10/31
package Repositories

import (
	"gin-example/app/models"
	"gin-example/utils/db"
)

type User models.Users

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
func (u *User) Create() (int64, error) {
	result := db.Mysql.Create(&u)
	return u.UserId, result.Error
}

// CheckAuth
//
// @Description: 验证用户是否存在
//
// @Author: Jade
//
// @Date: 2022-10-31 16:02:53
//
// @Receiver  u
//
// @Return  bool
// @Return  error
func (u *User) CheckAuth() (bool, error) {
	err := db.Mysql.Select("user_id").Where(User{Username: u.Username, Password: u.Password}).First(u).Error
	if err != nil {
		return false, err
	}

	if u.UserId > 0 {
		return true, nil
	}

	return false, nil
}

// GetUserIdByUserName
//
// @Description: 通过账号查找用户ID
//
// @Author: Jade
//
// @Date: 2022-10-31 16:06:38
//
// @Receiver  u
//
// @Param  userName
//
// @Return  int64
// @Return  error
func (u *User) GetUserIdByUserName(userName string) (int64, error) {
	err := db.Mysql.Select("user_id").Where(User{Username: userName}).First(u).Error
	if err != nil {
		return 0, err
	}

	return u.UserId, nil
}
