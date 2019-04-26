package main

import "fmt"

func main() {
	//slice
	x := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
