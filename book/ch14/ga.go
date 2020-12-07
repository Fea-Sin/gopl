/*
/*
@Time : 2020/12/7 11:50 上午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w)

	w = os.Stdout
	fmt.Printf("%T\n", w)

	w = new(bytes.Buffer)

	fmt.Printf("%T\n", w)

	//fmt.Println( w.Write([]byte("hello")) )
}
