package controllers

import (
	"context"
	"gin-example/init/config"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Index
//
// @Description: 首页
//
// @Author: Jade
//
// @Date: 2022-10-30 21:41:18
//
// @param c
func Index(c *gin.Context) {
	status := http.StatusOK
	message := "OK!"
	ttl := config.Cfg.JWT.ExpiresAt * time.Minute
	key := "token"
	rdb := db.Redis
	defer rdb.Close()
	ctx := context.Background()
	token := rdb.Get(ctx, key).Val()

	if token == "" {
		token, err := jwt.GenerateToken(123, "bac")
		if err != nil {
			log.Fatalln(err)
		}
		//ttl := conf.Cfg.JWT.ExpiresAt * time.Minute
		err = rdb.Set(ctx, key, token, ttl).Err()
		if err != nil {
			log.Fatalln(err)

		}

	}

	claims, err := jwt.ParseJwtToken(token)
	if err != nil {
		status = http.StatusInternalServerError
		message = err.Error()
	}

	data := gin.H{
		"status": status,
		"data": gin.H{
			"message":   message,
			"token":     token,
			"user_id":   jwt.GetUserId(claims),
			"user_name": claims.Username,
		},
	}
	clientToken, err := c.Cookie("token")
	if err != nil || clientToken == "" {
		//创建cookie
		//cooki_e = "NotSet"
		c.SetCookie("token", token, int(config.Cfg.JWT.ExpiresAt)*60, "/", "http://127.0.0.1:8080", false, true)
	}

	c.JSON(http.StatusOK, data)
}

func Hello(c *gin.Context) {
	rdb := db.Redis
	defer rdb.Close()
	ctx := context.Background()

	err := rdb.Set(ctx, "abc", "123", 0).Err()
	if err != nil {
		log.Fatalln(err)
	}
	str := rdb.Get(ctx, "abc").Val()

	data := gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"name":    "哈哈" + str,
			"age":     20,
			"hobby":   []string{"dad", "You trusted all proxies, this is NOT safe. We recommend you to set a value."},
			"message": "hello world! ",
		},
	}
	c.JSON(http.StatusOK, data)
}
