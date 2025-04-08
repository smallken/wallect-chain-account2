package main

import (
	"fmt"
	"math"
)

func printArea(s Shape) {
	fmt.Println("area:", s.Area())
}

func returnInter(radius float64) Shape {
	// 可以返回实现接口的结构体
	return Circle{Radius: radius}
}

type Cylinder struct {
	Redicus, Height float64
}

func (c Cylinder) Volume() float64 {
	return math.Pi * c.Redicus * c.Redicus * c.Height
}

func (c Cylinder) Area() float64 {
	return 2 * math.Pi * c.Redicus * (c.Redicus + c.Height)
}

func main() {
	c := Circle{Radius: 4}
	r := Retangle{Width: 4, Height: 5}
	printArea(c)
	printArea(r)
	fmt.Println(returnInter(6))
}
