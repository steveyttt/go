package main

import "fmt"

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Red",
		friends: map[string]int{
			"jimmy": 0414444444,
			"sally": 056555526,
		},
		favDrinks: []string{
			"scotch",
			"beer",
			"water",
		},
	}

	fmt.Println(s)

}
