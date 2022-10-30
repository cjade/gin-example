package jwt

import (
	"gin-example/app/services/user"
	"gin-example/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTConf struct {
	Secret    []byte
	ExpiresAt time.Duration
}

var jwtConf = &JWTConf{}

func init() {
	jwtConf.Secret = []byte(conf.Cfg.JWT.Secret)
	jwtConf.ExpiresAt = conf.Cfg.JWT.ExpiresAt * time.Minute
}

// GenerateToken
//
//	@Description: 生成token
//	@return string token
//	@return error
func GenerateToken(userName string, password string) (string, error) {
	mySigningKey := jwtConf.Secret
	_ = password
	var userId int64 = 123
	// Create the Claims
	claims := Claims{
		userId,
		userName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(jwtConf.ExpiresAt),
			),
			Issuer: "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	user.CacheToken(userId, ss)

	return ss, err
}

// ParseJwtToken
//
//	@Description: 解析token
//	@param tokenString
//	@return *Claims
//	@return error
func ParseJwtToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtConf.Secret, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}

// GetUserId
//
//	@Description: 获取用户ID
//	@param claims
//	@return int64
func GetUserId(claims *Claims) int64 {
	return claims.Id
}
