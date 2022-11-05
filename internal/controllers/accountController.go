package controllers

import (
	"context"
	"fmt"
	"gin-example/init/config"
	"gin-example/internal/services/user"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util/app"
	"gin-example/pkg/util/db"
	"gin-example/pkg/util/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UserID   int `uri:"id" binding:"required"`
	UserName string
}

var userService = user.Service{}

// Login
//
// @Description: 登录
//
// @Author: Jade
//
// @Date: 2022-10-31 00:16:29
//
// @Param  c
func Login(c *gin.Context) {
	appR := app.Gin{C: c}
	req := &user.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		log.Printf("%v", err.Error())
		appR.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	isExist := userService.Check(req)

	if !isExist {
		appR.Response(http.StatusBadRequest, e.AccountOrPasswordError, nil)
		return
	}

	token, err := jwt.GenerateToken(userService.UserId, req.UserName)
	if err != nil {
		appR.Response(http.StatusBadRequest, e.ERROR, nil)
		return
	}

	c.SetCookie("token", token, int(config.Cfg.JWT.ExpiresAt)*60, "/", "http://127.0.0.1:8080", false, true)
	appR.Response(http.StatusOK, 200,
		gin.H{
			"user_id": userService.UserId,
			"token":   token,
		},
	)

}

// Logout
// @Description: 退出登录
// @param c
func Logout(c *gin.Context) {
	ctx := context.Background()
	key := fmt.Sprintf(config.Cfg.CacheTokenKey, userService.GetUserId())
	rdb := db.Redis
	defer rdb.Close()

	if err := rdb.Del(ctx, key).Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1001,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success!",
	})
}

// Register
//
//	@Description: 注册
//	@param c
func Register(c *gin.Context) {
	appR := app.Gin{C: c}
	req := &userService.RegisterRequest

	if err := c.ShouldBind(req); err != nil {
		appR.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	// 账号是否已经存在
	isExistAccount := userService.IsExistAccount()
	if isExistAccount {
		appR.Response(http.StatusBadRequest, e.AccountExisting, nil)
		return
	}

	token, err := userService.Register()
	if !err {
		appR.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appR.Response(http.StatusOK, e.SUCCESS, gin.H{
		"user_id": userService.UserId,
		"token":   token,
	})
	return
}

// UserInfo
//
// @Description: 用户信息
// @param c
func UserInfo(c *gin.Context) {
	u := &User{}
	err := c.ShouldBindUri(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	userId := userService.GetUserId()
	c.JSON(http.StatusOK, gin.H{"uuid": userId})
}
