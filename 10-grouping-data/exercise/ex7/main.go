package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooo, James"}

	fmt.Println(x)
	fmt.Println(y)

	xxs := [][]string{x, y}
	fmt.Println(xxs)

	for i, xx := range xxs {
		fmt.Println("record:", i)
		for j, val := range xx {
			// fmt.Println(j, val)
			fmt.Printf("\t index position: %v \t value: %v \n", j, val)
		}
	}
}
