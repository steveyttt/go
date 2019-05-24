package main

import "fmt"

func main() {

	fmt.Println("Lets add some srings", mySum(1, 2, 3))
	fmt.Println("2 + 3", mySum(2, 3))
	fmt.Println("5 + 9", mySum(5, 9))
	fmt.Println("6 + 77", mySum(6, 77))

}
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum

}
