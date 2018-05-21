package interfaces

import (
	"fmt"
	"math"
)

func ExecuteShapes() {
	fmt.Println("Inside interfaces.ExecuteShapes")
	defer fmt.Println("Completed interfaces.ExecuteShapes")

	s := square{5}
	c := circle{4}
	r := rectangle{4, 5}

	fmt.Println("Square area:", s.area())
	fmt.Println("Circle area:", c.area())
	fmt.Println("Rectangle area:", r.area())

	printArea(s)  //value receiver and value type - works
	printArea(&c) //value receiver and pointer type - works
	printArea(&r) //pointer receiver and pointer type - works
	// printArea(r) //pointer receiver and value type - Boom
}

func printArea(s shape) {
	fmt.Println("Area of the shape is:", s.area())
}

type shape interface {
	area() float32
}

type square struct {
	side int
}

//value receiver
func (s square) area() float32 {
	return float32(s.side * s.side)
}

type circle struct {
	radius int
}

//value receier
func (c circle) area() float32 {
	return math.Pi * float32(c.radius*c.radius)
}

type rectangle struct {
	length int
	width  int
}

//pointer receiver
func (r *rectangle) area() float32 {
	return float32(r.length * r.width)
}
