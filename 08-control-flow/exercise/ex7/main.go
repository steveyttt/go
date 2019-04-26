package main

import "fmt"

func main() {
	x := "JamesBond"

	if x == "Money Penny" {
		fmt.Println(x)
	} else if x == "JamesBond" {
		fmt.Println("BOND James Bond")
	} else {
		fmt.Println("No Matches")
	}

}
