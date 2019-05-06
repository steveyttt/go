package main

//a POINTER receiver only works with values that are POINTERS
//a NON-POINTER receiver works with values that are POINTERS and NON-POINTERS

import (
	"fmt"
	"math"
)

//method sets are a set of methods that you can attach to a TYPE

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(&c)
}
