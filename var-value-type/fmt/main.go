package main

//https://godoc.org/fmt
import "fmt"

var y = 42

//character escapes are here - https://golang.org/ref/spec#Rune_literals
func main() {
	fmt.Println("2")
	fmt.Printf("%T\n", y)                // print TYPE
	fmt.Printf("%v\n", y)                // print TYPE
	fmt.Printf("%b\n", y)                // print binary
	fmt.Printf("%x\n", y)                // base 16
	fmt.Printf("%#x\n", y)               // unit value = hexadicimal
	fmt.Printf("%#x\t%b\t%x\n", y, y, y) // run it three times in one line

}
