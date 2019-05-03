package main

import "fmt"

func main() {
	defer foo() //defer causes this function to run once all other code in this function has executed
	bar()

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
