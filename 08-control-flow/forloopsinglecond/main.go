//https://golang.org/ref/spec#For_statements
package main

import "fmt"

func main() {
	x := 1
	y := 1

	//while x is less than true run
	//for statement with single condition
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("done")

	for {
		//if y is greater than 9 break from the for loop
		if y > 9 {
			break
		}

		fmt.Println(y)
		y++
	}

	fmt.Println("done")

}
