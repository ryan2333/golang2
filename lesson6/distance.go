package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y) //求两点之间的距离
}

func (p *Point) PDistance(q Point) float64 { //求p点到任意q点的距离
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	o := Point{10, 12}
	fmt.Println(Distance(p, q)) //5
	fmt.Println(p.PDistance(q)) //5
	fmt.Println(p.PDistance(o)) //13.45
}
