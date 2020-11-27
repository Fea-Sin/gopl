/*
/*
@Time : 2020/11/27 2:37 下午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import "fmt"

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x -1)
}

func main() {
	f(3)
}
