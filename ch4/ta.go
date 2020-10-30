/*
/*
@Time : 2020/10/30 4:48 下午
@Author : chengqunzhong
@File : ta
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
	"crypto/sha256"
)

func main() {
	for _, arg := range os.Args[1:] {
		c := sha256.Sum256([]byte(arg))
		fmt.Printf("动态生成HASH：%x\n", c)
	}
}

