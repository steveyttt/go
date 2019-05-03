//video 106 re-watch
//interfaces say that if you can use these methods you are of my type (You can run)
//this works as values can be of MULTIPLE types
// i.e. secret agent is also of tyoe human (as they can both run speak())
// i.e. person is also of type human (as they can both run speak())
//interfaces are created as a type

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

type hotdog int

//an interface says hey baby if you can run method speak, youre my TYPE :)
type human interface {
	speak()
}

//func (r receiver) identifier(parameters) return(s) () {code}
//func speak can only be invoked by struct type secret agent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- The secret agent speak")
}

//func speak can only be invoked by struct type person
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- The person speak")
}

//under the interface which takes many types create a switch / case statement to provide a different action based upon the provided type
// we say provide a parameter of h (type human) which is of type secretagent and person
func bar(h human) {
	switch h.(type) { //switch on the TYPE of the provided h parameter
	case person:
		fmt.Println("You are aaaaaaaaaa person", h.(person).first)
	case secretAgent:
		fmt.Println("You are aaaaaaaaaa secretAgent", h.(secretAgent).first)
	default:
		fmt.Println("unknown type")
	}
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
		first: "Doctor",
		last:  "Noooo",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	sa1.speak()
	sa2.speak()

	// this is polymorphism. a function which can take multiple types (the person type and the secret agent type)
	// interfaces allow values to be of more than one type
	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) //convert x to be an int
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
