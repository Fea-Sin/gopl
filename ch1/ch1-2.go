/*
/*
@Time : 2020/10/23 11:17 上午
@Author : chengqunzhong
@File : ch1-2
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}

