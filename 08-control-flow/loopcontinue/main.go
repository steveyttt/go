//print all even numbers from 0 - 100
package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue //If x divided by 2 has a remainder then continue to next loop
		}
		fmt.Println(x)
	}
}
