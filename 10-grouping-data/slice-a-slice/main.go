package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[1:]) //print from index 1
	y := x[1:3]        //print from index 1 up to but not including index 3
	fmt.Println(y)

	for i := 0; i < len(x); i++ {
		fmt.Println("Index", i, "of variable", x, "Holds value", x[i])
	}
}
