/*
/*
@Time : 2020/11/24 11:19 上午
@Author : chengqunzhong
@File : autoescape
@Software: GoLand
*/
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct{
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
