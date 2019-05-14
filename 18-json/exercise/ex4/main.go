package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{5, 6, 8, 2, 5, 23, 76, 12, 56, 89, 99}
	xs := []string{"mustache", "grapes", "panda", "curly", "monkey", "bored", "cheese", "faces", "clowns", "potato"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
