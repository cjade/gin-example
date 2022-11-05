// Package e
// @Description: 错误code信息
// @Author: Jade
// @Date: 2022/10/31
package e

var MsgFlags = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	AccountExisting:        "账号已存在",
	AccountOrPasswordError: "账号或密码已错误",
	InvalidParams:          "请求参数错误",
	RecordNotFound:         "记录不存在",
}

func GetMsg(code int) string {
	if s, ok := MsgFlags[code]; ok {
		return s
	}
	return MsgFlags[ERROR]
}
