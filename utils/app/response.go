package app

import (
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
// @Description:
//
// @Author: Jade
//
// @Date: 2022-10-30 21:45:37
//
// @Receiver  g
//
// @Param  httpCode
// @Param  errCode
// @Param  data
//
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		//Msg:  e.GetMsg(errCode),
		Msg:  "e.GetMsg(errCode)",
		Data: data,
	})
	return
}
