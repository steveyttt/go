package main

import "fmt"

//variadic parameters are functions which take many parameters

func main() {
	foo(66, 67)

	b := foo(2, 3, 4, 5, 6, 7)
	fmt.Println(b)
}

//basic function
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position", i, "we are now adding", v, "to the total", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
