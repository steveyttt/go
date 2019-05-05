//callback - passing a func as an argument
package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7}
	s := sum(ii...)
	fmt.Println("All numbers added equal", s)

	s2 := even(sum, ii...)
	fmt.Println("All even numbers added equal", s2)

	s3 := odd(sum, ii...)
	fmt.Println("All odd numbers added equal", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//func (r receiver) identifier (parameters) (return (s)) {function code}

//func even takes 2 params, one is a function (xi) and the second is a variadic parameter (vi)
// it returns an int
func even(f func(xi ...int) int, vi ...int) int {
	//create an empty slice
	var yi []int
	//range over the vi int provided to the function
	for _, v := range vi {
		//if v is even append to the slice
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	//return function sum using only the slice of even numbers created above
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
