package main

import (
	"fmt"
	"math"
)

type square struct {
	height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.height * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())

}

func main() {

	sq := square{
		height: 10,
	}

	ci := circle{
		radius: 22,
	}

	info(sq)
	info(ci)

}
