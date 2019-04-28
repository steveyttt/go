//this is all about  creating values of different types
//think of them as a data structure (OBJECT)
package main

import "fmt"

//Create a new type of person and have it be a struct
//composite data structure
//can contain values of differnt types i.e. int, string
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   38,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
