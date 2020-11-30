/*
/*
@Time : 2020/11/30 11:33 上午
@Author : chengqunzhong
@File : ga
@Software: GoLand
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}
type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
