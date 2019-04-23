// every go script needs a package main at the top
// you can only ever execute if it is called main
package main

//import FMT to the code
//fmt is a std lib package
import "fmt"

// Var - you can declare a variable outside of a fucntion body using the var keyword
// the shorthand declatation operator := will NOT work outside the function body
// it must exist within func main()
// best practice is to use short hand declarations at all times - keep variables at local scope

// variable can be called from anywhere in code
// package scope
var u = 44 //https://golang.org/ref/spec#Variable_declarations

// Declare a variable with identifier Z
// sets the VARIABLE as TYPE int (number)
// assigns the 0 value
// global vars done have to be called in a script
// local shorthand ones do
// https://golang.org/ref/spec#The_zero_value
var z int

//entry point for code
//when you leave func main, your code is done and stops running.
func main() {
	//You must use a variable if declared or code will fail
	// short decloratio operator :=
	x := 42 //declare a variable x: is declaring -- = 42 is assigning
	fmt.Println("howdy")
	fmt.Println(x) //call the println function in the fmt package
	x = 84         //reassign a variable (No need for : as it has been assigned already)
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
	fmt.Println(u)
	fmt.Println(z)
}
