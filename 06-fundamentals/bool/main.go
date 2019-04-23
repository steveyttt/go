package main

import (
	"fmt"
)

// Create a variable called x of a type bool
var x bool

// https://golang.org/ref/spec#Operators_and_punctuation
// https://golang.org/ref/spec#Boolean_types

func main() {
	a := 7
	b := 42
	//print x (by default x is false until proven true)
	fmt.Println(x)
	x = true
	//now prints true due to the above assignement
	fmt.Println(x)
	//conditional operator, if a equals b then answer true
	fmt.Println(a == b)
}
