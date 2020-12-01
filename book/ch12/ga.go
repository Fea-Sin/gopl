/*
/*
@Time : 2020/12/1 10:49 上午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint

	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Point.Y = 2
	fmt.Println(cp.Y)

	fmt.Println(cp.X)
}