package main

import "fmt"

func main() {
	s2 := foo()
	fmt.Println(s2)

	fmt.Println(bar()())

	y := bar()
	j := y()
	fmt.Println(j)

}

//func (r receiver) identifier (parameters) (return (s)) {function code}
//func foo returns a string
func foo() string {
	return "Hello world"
}

//func bar RETURNS a function which returns an INT
func bar() func() int {
	return func() int {
		return 451
	}
}
