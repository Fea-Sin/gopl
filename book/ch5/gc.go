/*
/*
@Time : 2020/11/23 6:00 下午
@Author : chengqunzhong
@File : gc
@Software: GoLand
*/
package main

import (
	"fmt"
	"gopl.io/book/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
