/*
/*
@Time : 2020/12/1 4:41 下午
@Author : chengqunzhong
@File : gd
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

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 60}

	distance := Point.Distance

	fmt.Println(distance(p, q))

	scale := (*Point).ScaleBy

	scale(&p, 30)
	fmt.Println(p)
}


