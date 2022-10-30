package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
	Server        ServerConf `yaml:"Server"`
	Mysql         MysqlConf  `yaml:"Mysql"`
	Redis         RedisConf  `yaml:"Redis"`
	Mongo         MongoConf  `yaml:"Mongo"`
	JWT           JWTConf    `yaml:"JWT"`
	CacheTokenKey string     `yaml:"CacheTokenKey"`
}

var Cfg = &Conf{}

func init() {

	path := "conf/config.yaml"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件.. error %v", err)
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("------- t:\n%v\n\n", Cfg)

	//d, err := yaml.Marshal(&Cfg)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

}
