// Package sf
// @Description: 雪花算法生成ID
// @Author: Jade
// @Date: 2022/10/31 21:56
package sf

import (
	"github.com/bwmarrin/snowflake"
)

// GenId
//
// @Description: 生成雪花ID
//
// @Author: Jade
//
// @Date: 2022-10-31 22:12:15
//
// @Return  int64
// @Return  error
func GenId() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	id := node.Generate().Int64()
	return id, nil
}
