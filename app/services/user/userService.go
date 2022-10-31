package user

import (
	"context"
	"fmt"
	"gin-example/app/Repositories"
	"gin-example/conf"
	"gin-example/utils/db"
	"time"
)

// Id user_id
var Id int64 = 0

func InitUserServer(userid int64) {
	Id = userid
}

type Auth struct {
	UserName string
	Password string
}

func (a *Auth) Check() (bool, error) {
	u := Repositories.User{Username: a.UserName, Password: a.Password}
	return u.CheckAuth()
}

// CacheToken
//
// @Description: 缓存token
// @param userID
// @param token
func CacheToken(userID int64, token string) {
	ctx := context.Background()
	key := fmt.Sprintf(conf.Cfg.CacheTokenKey, userID)
	ttl := conf.Cfg.JWT.ExpiresAt * time.Minute
	rdb := db.Redis
	defer rdb.Close()

	rdb.Set(ctx, key, token, ttl)
}
