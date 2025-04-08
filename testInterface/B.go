package main

import "math"

type Circle struct {
	Radius float64
}

type Retangle struct {
	Width, Height float64
}

type Shape3D interface {
	Shape
	Volume() float64
}

func (c Circle) Area() float64 {
	return c.Radius * math.Pi * c.Radius
}

func (r Retangle) Area() float64 {
	return r.Width * r.Height
}
