/*
/*
@Time : 2020/12/15 4:15 下午
@Author : chengqunzhong
@File : gb
@Software: GoLand
*/
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squarers := make(chan int)

	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}
