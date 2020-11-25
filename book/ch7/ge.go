/*
/*
@Time : 2020/11/25 4:23 下午
@Author : chengqunzhong
@File : ge
@Software: GoLand
*/
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
var depth int

func startElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("parse is down: %v", err)
	}

	forEachNode(doc, startElement, endElement)
}
