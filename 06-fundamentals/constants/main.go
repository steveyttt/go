//constants never change
package main

import "fmt"

const a = 42
const b = 42.78
const c = "Bobby"

const (
	d int     = 55
	e float64 = 66.66
	f string  = "james Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
