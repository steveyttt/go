package main

import "fmt"

//variadic parameters are functions which take many parameters

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum(xi...) //unfurl the slice before using (apply the numbers and not the slice)

	b := sum(2, 3, 4, 5, 6, 7)
	fmt.Println(b)
}

//basic function
func sum(x ...int) int {
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
