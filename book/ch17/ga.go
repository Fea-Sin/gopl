/*
/*
@Time : 2020/12/15 2:42 下午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 101; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}