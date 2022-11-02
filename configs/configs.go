// Package configs
// @Description:
// @Author: Jade
// @Date: 2022/11/2 23:35
package configs

import (
	"time"
)

type ServerConf struct {
	HostPort     int `yaml:"HostPort"` // 端口
	ReadTimeout  int `yaml:"ReadTimeout"`
	WriteTimeout int `yaml:"WriteTimeout"`
}

type MysqlConf struct {
	Host        string `yaml:"Host"`
	Port        int    `yaml:"Port"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	DbName      string `yaml:"DbName"`
	TablePrefix string `yaml:"TablePrefix"`
}

type RedisConf struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	DB       int    `yaml:"DB"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type MongoConf struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Auth     bool   `yaml:"Auth"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type JWTConf struct {
	Secret    string        `yaml:"Secret"`
	ExpiresAt time.Duration `yaml:"ExpiresAt"`
}

type Conf struct {
	Version       string     `yaml:"Version"`
	Debug         bool       `yaml:"Debug"`
	Server        ServerConf `yaml:"Server"`
	Mysql         MysqlConf  `yaml:"Mysql"`
	Redis         RedisConf  `yaml:"Redis"`
	Mongo         MongoConf  `yaml:"Mongo"`
	JWT           JWTConf    `yaml:"JWT"`
	CacheTokenKey string     `yaml:"CacheTokenKey"`
}
