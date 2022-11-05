// Package cache
// @Description:
// @Author: Jade
// @Date: 2022/11/3 23:48
package cache

import (
	"context"
	"fmt"
	"gin-example/init/config"
	"gin-example/pkg/util/db"
	"log"
	"time"
)

// CacheToken
//
// @Description: 缓存token
// @param userID
// @param token
func CacheToken(userID uint64, token string) {
	ctx := context.Background()
	key := fmt.Sprintf(config.Cfg.CacheTokenKey, userID)
	log.Printf("cache token key :%s", key)

	ttl := config.Cfg.JWT.ExpiresAt * time.Minute
	rdb := db.Redis

	err := rdb.Set(ctx, key, token, ttl).Err()
	if err != nil {
		log.Printf("cache token key error :%s", err.Error())
	}
}
