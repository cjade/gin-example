package user

import (
	"gin-example/internal/models"
	"gin-example/internal/repositories"
	"gin-example/pkg/jwt"
	"gin-example/pkg/util/hash"
	"log"
)

type RegisterRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Service struct {
	UserId uint64
	RegisterRequest
}

var repoUser repositories.UsersRepository

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
func (s *Service) Check(req *LoginRequest) bool {
	repoUser.Fields = []string{"user_id", "password"}
	user := repoUser.GetUserByAccount(req.UserName)
	s.UserId = user.UserId

	log.Printf("password : [%s] hashPassword : [%s]", req.Password, user.Password)
	return hash.BcryptCheck(req.Password, user.Password)
}

func (s *Service) Register() (string, bool) {
	u := models.Users{}
	u.UserName = s.UserName
	u.Password = hash.BcryptHash(s.Password)
	userId, err := repoUser.Create(u)
	if err != nil {
		return "", false
	}
	s.UserId = userId
	token, err := jwt.GenerateToken(userId, u.UserName)
	if err != nil {
		return "", false
	}

	return token, true
}

// IsExistAccount
//
// @Description:  账号是否已存在
//
// @Author: Jade
//
// @Date: 2022-11-03 18:39:43
//
// @Receiver  s
//
// @Return  bool  true - 存在 ,false - 存在
func (s *Service) IsExistAccount() bool {
	repoUser.Fields = []string{"user_id"}
	users := repoUser.GetUserByAccount(s.UserName)

	if users.UserId > 0 {
		return true
	}
	return false
}

func (s *Service) SetUserId(userid uint64) {
	s.UserId = userid
}
func (s *Service) GetUserId() uint64 {
	return s.UserId
}
