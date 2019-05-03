//store a function as an expression
package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("my first func expression")
	}

	g := func(x int) {
		fmt.Println("lucky number", x)
	}

	f()
	g(1984)
}
