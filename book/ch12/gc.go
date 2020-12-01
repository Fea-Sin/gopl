/*
/*
@Time : 2020/12/1 4:28 下午
@Author : chengqunzhong
@File : gc
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
	q := Point{4, 6}

	distanceFromP := p.Distance

	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)
}
