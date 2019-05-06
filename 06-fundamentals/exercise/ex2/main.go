package main

import "fmt"

func main() {
	a := (42 == 42) //equal
	b := (40 <= 39) //less than or equal
	c := (42 >= 42) // greater than or equal
	d := (40 != 40) //not equal
	e := (40 < 42)  //less than
	f := 45 > 42    //greater than

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
