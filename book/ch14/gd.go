/*
/*
@Time : 2020/12/9 10:39 上午
@Author : chengqunzhong
@File : gd
@Software: GoLand
*/
package main

import "fmt"

type Errno int

var errors = [...]string {
	1: "operation not permitted",
	2: "no such file or directory",
	3: "no such process",
}

func (e Errno) Error() string {
	if 0 < e && int(e) < len(errors) {
		return errors[e]
	}
	return fmt.Sprintf("errno %d", e)
}

func main() {
	err := Errno(2)

	fmt.Println(err)
}


