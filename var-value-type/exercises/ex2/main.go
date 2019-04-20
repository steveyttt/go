package main

import "fmt"

// These are zero value variables
var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	fmt.Println("These are known as zero value variables ^^")
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
