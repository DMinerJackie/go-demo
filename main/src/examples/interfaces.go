package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rects1 struct {
	width, height float64
}

type cicle struct {
	radius float64
}

func (r rects1) area() float64 {
	return r.width * r.height
}

func (r rects1) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c cicle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cicle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rects1{width: 3, height: 4}
	c := cicle{radius: 5}

	measure(r)
	measure(c)
}
