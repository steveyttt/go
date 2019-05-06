// func receiver identifier params code
package main

import "fmt"

func main() {

	x := foo()
	yint, ystring := bar()

	fmt.Println(x)
	fmt.Println(yint, ystring)

}

func foo() int {
	return 1
}
func bar() (int, string) {
	return 0, "Welcome to the party"
}
