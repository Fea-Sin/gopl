/*
/*
@Time : 2020/11/20 11:01 上午
@Author : chengqunzhong
@File : gb
@Software: GoLand
*/
package main

import (
	"sort"
	"fmt"
)

func main() {
	var names []string
	var ages = map[string]int {}

	ages["alice"] = 31
	ages["charlie"] = 34
	ages["bob"] = 3

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
