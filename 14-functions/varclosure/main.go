//enclosing a variable in some scope
//var closure, var scope
package main

import "fmt"

//this is a package wide scope
var x int

func main() {

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{
		y := 42
		fmt.Println(y)
	}

	//the below fails as y is only in scope for the above curly brackets
	// y++

	//here you see x is the variable in the function called incrementor
	//by creating 2 var's for the function it creates 2 instances of the x variable
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	x++
}

func incrementor() func() int {
	var x int //this var x exists all the way through this function
	return func() int {
		x++
		return x
	}

}
