// Package main
// @Description:
// @Author: Jade
// @Date: 2022/11/2 16:41
package main

import (
	"flag"
	"fmt"
)

var name = flag.String("flagname", "sd", "help message for flagname")

func main() {
	flag.Parse()
	fmt.Println(*name)
	aa()
}
