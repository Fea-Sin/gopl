/*
/*
@Time : 2020/11/19 2:27 下午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	ages := map[string]int {}
	ages["alice"] = 32
	ages["charlie"] = 34

	fmt.Println(ages["alice"])

	//delete(ages, "alice")

	ages["bob"] = ages["bob"] + 1

	fmt.Println(ages["bob"])

	for name, age := range ages {
		fmt.Printf("%s\t%d", name, age)
		fmt.Println()
	}
}