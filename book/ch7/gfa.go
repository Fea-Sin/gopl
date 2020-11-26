/*
/*
@Time : 2020/11/26 10:46 上午
@Author : chengqunzhong
@File : gfa
@Software: GoLand
*/
package main

import "fmt"

// squares 返回一个匿名函数
// 该匿名函数每次调用时都会返回下一个数的平方
func squares() func() int  {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	// 1. 匿名函数使用
	//fmt.Println( strings.Map( func(r rune) rune { return r +1}, "HAL-9000" ) )

	// 2. 闭包
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
