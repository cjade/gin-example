package routes

import (
	"gin-example/internal/controllers"
	"gin-example/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.New()
	r.GET("/index", controllers.Index)
	r.GET("/get_book", controllers.GetBook)
	r.GET("/hello", controllers.Hello)
	// 登录
	r.POST("/login", controllers.Login)
	// 注册
	r.POST("/register", controllers.Register)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth())
	{
		// 用户信息
		authorized.GET("/userinfo/:id", controllers.UserInfo)
		// 退出登录
		authorized.POST("/logout", controllers.Logout)
	}

	return r
}
