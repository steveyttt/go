package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) return(s) () {code}

// the secret agent func can only be called by a RECEIVER of type secretAgent
// this is called a method
// methods can only be called from type struct
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"MoneyPenny",
		},
		ltk: true,
	}

	p1 := person{
		"Mrs",
		"M",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	sa1.speak()
	sa2.speak()

}
