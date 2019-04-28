//this is all about  cfreating values of different types
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

//referencing the type person inside a new type called secretAgent
type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "collision", // works but ugly. Best not to have double key names.
		ltk:   true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
