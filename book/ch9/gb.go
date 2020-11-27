/*
/*
@Time : 2020/11/27 11:41 上午
@Author : chengqunzhong
@File : gb
@Software: GoLand
*/
package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)

	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func main() {
	bigSlowOperation()
}
