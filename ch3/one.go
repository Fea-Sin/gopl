/*
/*
@Time : 2020/10/26 5:01 下午
@Author : chengqunzhong
@File : one
@Software: GoLand
*/
package main

import "fmt"

func main() {
	ascii := 'a'
	unicode := '中'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

}
