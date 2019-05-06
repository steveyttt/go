package main

import "fmt"

func main() {
	foo()
	defer bar()
	too()

}

func foo() {
	fmt.Println("I am 1")
}

func bar() {
	fmt.Println("I am 2")
}

func too() {
	fmt.Println("I am 3")
}
