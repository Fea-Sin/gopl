/*
/*
@Time : 2020/11/11 1:43 下午
@Author : chengqunzhong
@File : trans
@Software: GoLand
*/
package main

import (
	"fmt"

	"github.com/mind1949/googletrans"
	"golang.org/x/text/language"
)

var text = "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

func main() {
	params := googletrans.TranslateParams{
		Src:  "auto",
		Dest: language.SimplifiedChinese.String(),
		Text: text,
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		panic(err)
	}
	fmt.Println("翻译：")
	fmt.Println(translated.Text)
}
