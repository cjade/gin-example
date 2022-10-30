package models

import "gin-example/utils/db"

type Users struct {
	UserId   int64  `json:"user_id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth
// Description: 验证用户密码是否正确
// @param username
// @param password
// @return bool
// @return error
func CheckAuth(username, password string) (bool, error) {
	var user Users
	err := db.Mysql.Select("user_id").Where(Users{Username: username, Password: password}).First(&user).Error
	if err != nil {
		return false, err
	}

	if user.UserId > 0 {
		return true, nil
	}

	return false, nil
}
