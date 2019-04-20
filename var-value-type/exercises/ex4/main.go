package main

import "fmt"

// These are zero value variables
type steveyt int

var x steveyt

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
