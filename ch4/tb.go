/*
/*
@Time : 2020/10/30 6:01 下午
@Author : chengqunzhong
@File : tb
@Software: GoLand
*/
package main

import (
	"fmt"
)

var abc = []string{"a", "a", "b", "b", "b", "c", "d", "j", "r", "r", "r"}

func king(src []string)  {
	var out []string
	var k = ""

	for _, v := range src {
		if k != v {
			out = append(out, v)
			k = v
		}
	}
	fmt.Println(out)
}

func main() {
	king(abc)
}

