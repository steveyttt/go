package main

import "fmt"

const (
	a = iota + 2014
	b = iota + 2014
	c = iota + 2014
	d = iota + 2014
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
