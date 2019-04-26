package main

import "fmt"

func main() {
	//slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[0:5])  //could also do x[:5]
	fmt.Println(x[5:10]) //could also do x[5:]
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
