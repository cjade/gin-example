// Package env
// @Description:
// @Author: Jade
// @Date: 2022/11/2 22:20
package env

import (
	"flag"
	"log"
	"strings"
)

var environment = flag.String("env", "", "请输入当前运行的环境:\n"+
	"dev:  开发环境 \n"+
	"test: 测试环境 \n"+
	"rel:  预发布环境 \n"+
	"prod: 正式环境 \n"+
	"")

func init() {
	flag.Parse()

	switch strings.ToLower(*environment) {
	case "dev":
	case "test":
	case "rel":
	case "prod":
	default:
		log.Fatalln("Warning: '-env' cannot be found, or it is illegal.")
	}

}

func Value() string {
	return *environment
}

func IsDev() bool {
	return *environment == "dev"
}

func IsTest() bool {
	return *environment == "test"
}
func IsRel() bool {
	return *environment == "rel"
}
func IsProd() bool {
	return *environment == "prod"
}
