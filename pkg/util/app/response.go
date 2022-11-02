package app

import (
	"gin-example/pkg/util/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response
//
// @Description: 响应json
//
// @Author: Jade
//
// @Date: 2022-10-31 00:47:15
//
// @Receiver  g
//
// @Param  httpCode
// @Param  errCode
// @Param  data
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
