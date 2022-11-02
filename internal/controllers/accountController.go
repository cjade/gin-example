package controllers

import (
	"context"
	"fmt"
	"gin-example/init/config"
	"gin-example/internal/repositories"
	"gin-example/internal/services/user"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util/app"
	"gin-example/pkg/util/db"
	"gin-example/pkg/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UserID   int `uri:"id" binding:"required"`
	UserName string
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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
	req := &LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		appR.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	authServer := user.Auth{}
	isExist, err := authServer.Check(req.UserName, req.Password)
	if err != nil {
		appR.Response(http.StatusBadRequest, e.RecordNotFound, nil)
		return
	}

	if !isExist {
		appR.Response(http.StatusBadRequest, e.RecordNotFound, nil)
		return
	}

	token, err := jwt.GenerateToken(req.UserName, req.Password)
	if err != nil {
		appR.Response(http.StatusBadRequest, 10003, nil)
		return
	}

	c.SetCookie("token", token, int(config.Cfg.JWT.ExpiresAt)*60, "/", "http://127.0.0.1:8080", false, true)
	appR.Response(http.StatusOK, 200,
		gin.H{
			"token": token,
		},
	)

}

// Logout
// @Description: 退出登录
// @param c
func Logout(c *gin.Context) {
	ctx := context.Background()
	key := fmt.Sprintf(config.Cfg.CacheTokenKey, user.UserId)
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
	b := repositories.Book{}
	b.BookName = "abd"
	bid, _ := b.Create()

	appR.Response(http.StatusOK, 200,
		gin.H{
			"bid":    bid,
			"BookId": b.BookId,
		},
	)
	return

	//if e

	userName := c.PostForm("user_name")
	password := c.PostForm("password")

	//hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	//if err != nil {
	//	log.Println(err)
	//}
	//encodePWD := string(hash)
	c.JSON(http.StatusOK, gin.H{
		"userName": userName,
		"password": password,
		//"encodePWD": encodePWD,
	})
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

	userId := user.UserId
	c.JSON(http.StatusOK, gin.H{"uuid": userId})
}
