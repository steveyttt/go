package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("This shouldnt print mate")
	case true:
		fmt.Println("This should print mate")
	}
}
