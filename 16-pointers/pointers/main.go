package main

//& is the address in memory
//* is the pointer to the value in memory

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //This is the value A's address in memory (Not the value, just the address)

	fmt.Printf("%T\n", a)  //This will show as an int
	fmt.Printf("%T\n", &a) //This will show as *int (A pointer to an int) //*int is it's own type and NOT an int

	b := &a
	fmt.Println(b)   //will show the address of a
	fmt.Println(*b)  //will show the value in the address (a pointer)
	fmt.Println(*&b) //will show 42

	*b = 43        //set the value in that memory address to be 43
	fmt.Println(a) //query the in-memory space, it should now be 43

}
