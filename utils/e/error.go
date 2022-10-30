package e

import "fmt"

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.Code, e.Msg)
}

func GetCode(err error) int {
	if e, ok := err.(MyError); ok {
		return e.Code
	}
	return -1
}
func GetMsg(err error) string {
	if e, ok := err.(MyError); ok {
		return e.Msg
	}
	return ""
}

func GetError(err error) error {
	if e, ok := err.(MyError); ok {
		return e
	}
	return nil
}
