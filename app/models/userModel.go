package models

type Users struct {
	UserId   int64  `json:"user_id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
