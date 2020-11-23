/*
/*
@Time : 2020/11/23 11:47 上午
@Author : chengqunzhong
@File : gb
@Software: GoLand
*/
package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
}