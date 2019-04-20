package main

import "fmt"

// These are zero value variables
type steveyt int

var x steveyt
var y int

func main() {
	//print x (it will be a zero value)
	fmt.Println(x)
	//print type of x, it is steveyt
	fmt.Printf("%T\n", x)
	//set x to be 42
	x = 42
	fmt.Println(x)
	// set y to be value x (42) with a TYPE of INT and not STEVEYT
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
