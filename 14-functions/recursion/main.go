//can use a loop insteas of a recursion
//recursion is a function calling itself
package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)

	n2 := loopfactorial(5)
	fmt.Println(n2)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopfactorial(x int) int {
	total := 1
	for ; x > 0; x-- {
		total *= x
	}
	return total
}
