/*
/*
@Time : 2020/12/16 4:36 下午
@Author : chengqunzhong
@File : gc
@Software: GoLand
*/
package main

import "fmt"

// buffer是1，channel会交替出现或空或满
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch<- i:
		}
	}
}

// buffer不是1，select会像抛硬币一样随机执行case
//func main() {
//	ch := make(chan int, 3)
//	for i := 0; i < 10; i++ {
//		select {
//		case x := <-ch:
//			fmt.Println(x)
//		case ch <- i:
//		}
//	}
//}
