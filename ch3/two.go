/*
/*
@Time : 2020/10/27 11:19 上午
@Author : chengqunzhong
@File : two
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "hello, 世界"
	n := 0
	t := "hello, 世界"
	k := "123"
	a, _ := strconv.ParseInt(k, 10, 64)
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		n++
	}
	fmt.Println("字符长度：", n)
	fmt.Println("是否全等：", s == t)
	fmt.Println(a, k)
}

