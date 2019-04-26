package main

import "fmt"

//composite data type or compound data type is any data type which can be constructed in a program using the programme langugaes primitive data types
//struct is a type which allows you to combine different data types into one list
//aggregate data type is for values of the same type i.e. array
func main() {
	//composite literal - create different values of a type
	// x := type{values}
	//group together values OF THE SAME TYPE
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(x)
	fmt.Println(x[0]) //query by index position
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(cap(x))

	//for index and value in the range of x print
	for i, v := range x {
		fmt.Println(i, v)
	}
}
