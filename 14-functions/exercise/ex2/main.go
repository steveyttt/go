package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4}
	y := foo(x...)
	z := bar(x)
	fmt.Println(y)
	fmt.Println(z)

}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total

}
