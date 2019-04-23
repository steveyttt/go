package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	a := 2
	//%d = decimal
	//%b = binary
	fmt.Printf("%d\t\t%b\n", a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b)

	// kb := 1024
	// mb := kb * 1024
	// gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
