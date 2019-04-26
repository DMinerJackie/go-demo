package main

import "fmt"

type rects struct {
	width, height int
}

func (r *rects) area() int {
	return r.width * r.height
}

func (r rects) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rects{width: 10, height: 20}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
