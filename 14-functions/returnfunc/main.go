package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	// fmt.Println(x)
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

}

//func (r receiver) identifier (parameters) (return (s)) {function code}
//func foo returns a string
func foo() string {
	s := "Hello world"
	return s
}

//func bar RETURNS a function which returns an INT
func bar() func() int {
	return func() int {
		return 451
	}
}
