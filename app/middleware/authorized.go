package middleware

import (
	"gin-example/app/services/user"
	jwt2 "gin-example/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = http.StatusOK
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		claims, err := jwt2.ParseJwtToken(token)

		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = 20002
			default:
				code = 20001
			}

		}
		UserId := jwt2.GetUserId(claims)
		user.InitUserServer(UserId)

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
