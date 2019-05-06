package main

import "fmt"

func main() {

	x := bar()
	i := x()
	fmt.Println(i)

}

func bar() func() int {
	return func() int {
		return 42
	}
}
