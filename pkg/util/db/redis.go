package db

import (
	"fmt"
	"gin-example/configs"
	"github.com/go-redis/redis/v8"
)

//type RedisUtil struct {
//	client redis.Conn
//}

var Redis *redis.Client

// RedisClient 全局变量, 外部使用utils.RedisClient来访问
//var RedisClient RedisUtil

func init() {
	address := fmt.Sprintf("%s:%d", configs.Cfg.Redis.Host, configs.Cfg.Redis.Port)

	rdb := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     address,
		Username: configs.Cfg.Redis.Username,
		Password: configs.Cfg.Redis.Password,
		DB:       configs.Cfg.Redis.DB,
	})

	Redis = rdb
}

//
//type RedisErr struct {
//	message string
//}
//
//// 方法名字是Error()
//func (e RedisErr) Error() string {
//	return fmt.Sprintf("%v ", e.message)
//}

//// SetStr 设置string
//func (db *RedisUtil) SetStr(key string, value string) error {
//	_, err := db.client.Do("set", key, value)
//	return err
//}
//
//// GetStr 获取string
//func (db *RedisUtil) GetStr(key string) (string, error) {
//	value, err := db.client.Do("get", key)
//	return string(value.([]byte)), err
//}
//
//// SetExpire 设置key过期时间
//func (db *RedisUtil) SetExpire(key string, expireSecond int) error {
//	_, err := db.client.Do("EXPIRE", key, expireSecond)
//	return err
//}
