//integers hard numbers - 0,1,2,3
//floats - 1.01, 2.02 (has fraction)
//https://golang.org/ref/spec#Numeric_types
package main

import "fmt"

//int can be positive and negative
//uint can only be positive
//byte = uint8 (0-255)
//rune = int32 (-2147483648 to 2147483647)
// the number at the after int and uint dictates the number of bytes to assign the variable
var x int
var y float64
var z int8 = 127   // int 8 can go up to 127, if you set this to 128 it will fail
var zz int16 = 256 // int 16 can go up to 32767

func main() {
	x := 42
	y := 42.44
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(zz)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
