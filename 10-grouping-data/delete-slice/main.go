//https://golang.org/ref/spec#Appending_and_copying_slices
package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42} //create a slice
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014) //append to var x these additional numbers
	fmt.Println(x)

	y := []int{234, 467, 789, 1099}
	x = append(x, y...) //merge var x and y (three dots ... are needed) //... unfurls the data and only provides the values and not the var type
	fmt.Println(x)

	x = append(x[:2], x[4:]...) //use append to remove values between index 2 and up to but not including index 4
	fmt.Println(x)

	x = append(x[:0], x[1:]...) //use append to remove first vale inside x slice
	fmt.Println(x)

}
