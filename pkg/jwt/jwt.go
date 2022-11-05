package jwt

import (
	"gin-example/init/config"
	"gin-example/internal/services/cache"
	"gin-example/pkg/env"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTConf struct {
	Secret    []byte
	ExpiresAt time.Duration
}

var jwtConf = &JWTConf{}

func init() {
	jwtConf.Secret = []byte(config.Cfg.JWT.Secret)
	jwtConf.ExpiresAt = config.Cfg.JWT.ExpiresAt * time.Minute
}

// GenerateToken
//
//	@Description: 生成token
//	@return string token
//	@return error
func GenerateToken(userId uint64, userName string) (string, error) {
	mySigningKey := jwtConf.Secret
	issuer := env.Value()
	// Create the Claims
	claims := Claims{
		userId,
		userName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(jwtConf.ExpiresAt),
			),
			Issuer: issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString(mySigningKey)
	if err == nil {
		cache.CacheToken(userId, strToken)
	}
	return strToken, err
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
func GetUserId(claims *Claims) uint64 {
	return claims.Id
}
