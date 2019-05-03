package main

import "fmt"

//func with no name

func main() {

	func(x int, y string) {
		fmt.Println("You are", 42, "years old", y)
	}(42, "Stevey")

	func() {
		fmt.Println("You are anonymous")
	}()

}
