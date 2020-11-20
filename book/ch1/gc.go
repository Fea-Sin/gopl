/*
/*
@Time : 2020/11/20 4:27 下午
@Author : chengqunzhong
@File : gc
@Software: GoLand
*/
package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	hello := map[string]int {
		"one": 1,
		"two": 2,
	}
	world := map[string]int {
		"one": 1,
		"two": 2,
		"three": 3,
	}

	// go 语法真是简洁·严苛
	//test := map[string]string {
	//	"one": "1",
	//	"two": "2",
	//}

	fmt.Printf("两个map是否相等%t\n", equal(hello, world))
}