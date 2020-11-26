/*
/*
@Time : 2020/11/26 4:23 下午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int) {}

func main() {
	// 参数已经是切片
	values := []int{1, 2, 3, 4}

	fmt.Println(sum())
	fmt.Println(sum(4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	fmt.Println(sum(values...))

	// 可变参数类型
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}
