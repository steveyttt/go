package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.last)
	fmt.Println(p.age)

}

func main() {

	p1 := person{
		first: "Steven",
		last:  "Tyson",
		age:   32,
	}

	p1.speak()

}
