// Package config
// @Description:
// @Author: Jade
// @Date: 2022/11/2 23:31
package config

import (
	"gin-example/configs"
	"gin-example/pkg/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg = &configs.Conf{}

func init() {
	e := env.Value()

	path := "configs/config." + e + ".yaml"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件.. error %v", err)
	}
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
