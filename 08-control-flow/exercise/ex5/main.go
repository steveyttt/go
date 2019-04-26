package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i, "When divided by 4 leaves modulus (remainder) of", i%4)
	}

}
