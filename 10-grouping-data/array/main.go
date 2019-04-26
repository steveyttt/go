package main

import "fmt"

func main() {
	// https://golang.org/ref/spec#Array_types
	//dont use arrays use slices
	var x [5]int //create an array with 5 entry points and type int
	fmt.Println(x)
	x[3] = 42 //set the third index to be 42 (from index 0 thats point 4)
	fmt.Println(x)
	fmt.Println(len(x))
}
