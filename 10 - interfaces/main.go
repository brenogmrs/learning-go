package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Printf("The area of this shape is %0.2f \n", s.area())
}

type square struct {
	height float64
	width  float64
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	fmt.Println("Interfaces")

	r := rectangle{10, 15}
	printArea(r)

	c := circle{10}
	printArea(c)
}
