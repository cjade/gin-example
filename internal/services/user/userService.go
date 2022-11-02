package user

import (
	"context"
	"fmt"
	"gin-example/configs"
	"gin-example/internal/repositories"
	"gin-example/pkg/util/db"
	"time"
)

// Id user_id
var UserId int64 = 0

func InitUserServer(userid int64) {
	UserId = userid
}

type Auth struct{}

// Check
//
// @Description:
//
// @Author: Jade
//
// @Date: 2022-11-01 14:50:32
//
// @Receiver  a
//
// @Param  userName
// @Param  password
//
// @Return  bool
// @Return  error
func (a *Auth) Check(userName, password string) (bool, error) {
	u := &repositories.Users{}
	u.UserName = userName
	u.Password = password
	//u := Repositories.User{UserName: a.UserName, Password: a.Password}
	return u.CheckAuth()
}

// CacheToken
//
// @Description: 缓存token
// @param userID
// @param token
func CacheToken(userID int64, token string) {
	ctx := context.Background()
	key := fmt.Sprintf(configs.Cfg.CacheTokenKey, userID)
	ttl := configs.Cfg.JWT.ExpiresAt * time.Minute
	rdb := db.Redis
	defer rdb.Close()

	rdb.Set(ctx, key, token, ttl)
}
