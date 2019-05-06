package main

import "fmt"

var a int = 42

func main() {
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)
	//Ths means shift all bits left by 1
	//this essentially doubles 42 to be 84
	b := a << 1
	fmt.Printf("%d\n%b\n%#x\n", b, b, b)
}
