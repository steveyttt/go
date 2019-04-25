package main

import "fmt"

func main() {
	fmt.Println(true && true)  //true AND true make true
	fmt.Println(true && false) //true AND false together make false - both need to be true
	fmt.Println(true || true)  //|| is OR, either can be true
	fmt.Println(true || false) //|| is OR, either can be true
	fmt.Println(!true)         //not true (false)
	fmt.Println(!false)        //not false (true)
}
