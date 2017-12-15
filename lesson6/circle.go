package main

import (
	"fmt"
	"math"
)

// var PI = 3.14
type Areaer interface { //声明interface结构体，调用接口需实现Area方法，返回float64类型
	Area() float64
}

type Circle struct {
	R float64
}

type Squal struct {
	L, H float64
}

func (c Circle) Area() float64 { //为Circle圆形实现Area方法 ，求圆形面积
	// return PI * c.R * c.R
	return math.Pi * math.Pow(c.R, 2)
}

func (c Squal) Area() float64 { //为Squal正方形实现Area方法，求正方形或长方形面积
	return c.L * c.H
}

func main() {
	var a Areaer

	// r := Circle{3}
	// a = r
	a = Circle{3}
	fmt.Println(a.Area()) //调用接口求圆面积

	// c := Squal{4, 5}
	// a = c
	a = Squal{4, 5}
	fmt.Println(a.Area()) //调用接口求正方形或长方形面积

}
