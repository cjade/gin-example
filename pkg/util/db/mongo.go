package db

import (
	"context"
	"fmt"
	"gin-example/init/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Mongo *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	auth := ""
	if config.Cfg.Mongo.Auth {
		auth = fmt.Sprintf("%s:%s@", config.Cfg.Mongo.Username, config.Cfg.Mongo.Password)
	}

	uri := fmt.Sprintf("mongodb://%s%s:%d", auth, config.Cfg.Mongo.Host, config.Cfg.Mongo.Port)
	opts := options.Client().ApplyURI(uri)

	// 链接数据库
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatalln("mongo 链接数据库失败  error : ", err.Error())
	}

	// 判断服务是否可用
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln("mongo 服务不可用 error : ", err.Error())
	}
	Mongo = client
}
