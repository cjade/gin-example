package hash

import "golang.org/x/crypto/bcrypt"

// BcryptHash
//
//	@Description: 使用 bcrypt 对密码进行加密
//	@param password
//	@return string
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck
//
// @Description: 对比明文密码和数据库的哈希值
// @param password
// @param hash
// @return bool
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
