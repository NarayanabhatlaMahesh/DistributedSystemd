package main

import (
	"fmt"
)

type Rectangle struct {
	length     int
	width      int
	area       float64
	perimeter  float64
}

func (r *Rectangle) calculateArea() {
	r.area = float64(r.length * r.width)
}

func (r *Rectangle) calculatePerimeter() {
	r.perimeter = 2 * float64(r.length+r.width)
}

func main() {
	rectangle := Rectangle{length: 5, width: 10}
	rectangle.calculateArea()
	rectangle.calculatePerimeter()
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", rectangle.area, rectangle.perimeter)
}
