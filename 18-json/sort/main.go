package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{4, 5, 6, 7, 8, 34, 45, 56, 67, 34, 67, 78, 12, 3, 34, 34, 5}
	xs := []string{"Hello", "I'm here", "And i am", "Gonna", "Bake you a cake"}

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("---------")

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

}
