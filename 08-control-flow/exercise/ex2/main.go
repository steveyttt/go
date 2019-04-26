package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for x := 0; x <= 2; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
