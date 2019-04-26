//https://golang.org/doc/effective_go.html#allocation_make
package main

import "fmt"

func main() {
	x := make([]int, 10, 50) //create an array of maximum 50 ints and creates a slice structure with a length of 10 (Length and capacity) - this allows the slice to grow by 40 more values until the arra needs to expand
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //number os spots inside the array we can use
	x[4] = 44
	x[5] = 999
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //number os spots inside the array we can use

}
