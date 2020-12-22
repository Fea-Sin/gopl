/*
/*
@Time : 2020/12/22 11:20 上午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)
//
//func main() {
//	var x, y int
//	fmt.Println("测试输出----")
//	xsema := make(chan struct{})
//	ysema := make(chan struct{})
//
//	go func() {
//		x = 1
//		fmt.Println("y:", y)
//		xsema <- struct{}{}
//	}()
//
//	go func() {
//		y = 1
//		fmt.Println("x:", x)
//		ysema <- struct{}{}
//	}()
//
//	<-xsema
//	<-ysema
//}

func main() {
	var x, y int
	fmt.Println("测试输出----")

	for {
		go func() {
			x = 1
			fmt.Print("y:", y, " ")
		}()

		go func() {
			y = 1
			fmt.Print("x:", x, " ")
		}()
	}

	time.Sleep(100 * time.Millisecond)

}


