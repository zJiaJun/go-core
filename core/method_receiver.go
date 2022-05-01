package core

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func MethodReceiverExample() {
	defer fmt.Println(Separator)
	r := rectangle{12, 2}
	c := circle{3}
	fmt.Println("MethodReceiverExample rectangle area", r.area())
	fmt.Println("MethodReceiverExample circle area", c.area())
}
