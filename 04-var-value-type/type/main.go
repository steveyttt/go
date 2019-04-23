package main

import "fmt"

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string (Other types are)
// Trying to reassign that value as a int will fail
// https://golang.org/ref/spec#Types
var z string = "Shaken, not stirred"

//set variable a as type INT
var a int

//create a new TYPE called HOTDOG which is part of the underlying type INT
//note you cannot set an int type variable to a hotdog type variable, they are different types and it will not work
//UNLESS you convert the variable
// https://golang.org/ref/spec#Conversions
type hotdog int

//set var b to be of TYPE hotdog
var b hotdog

// Use var for zero value variables
// back ticks just include everything inside the string
var zz string = `James said - "Shaken, not stirred"`

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", b)

	fmt.Println(z)
	fmt.Println(zz)
	fmt.Printf("%T \n", z)

	//conversion - a not equals variable b (with b converted from TYPE hotdog to INT)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
