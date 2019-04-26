//https://golang.org/ref/spec#Appending_and_copying_slices
package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42} //create a slice - slices are built on top of arrays - if you extend a slice it will need compute power to ditch the underlying array and rebuild for the new size.
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014) //append to var x these additional numbers
	fmt.Println(x)

	y := []int{234, 467, 789, 1099}
	x = append(x, y...) //merge var x and y (three dots ... are needed)
	fmt.Println(x)

}
