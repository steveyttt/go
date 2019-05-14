package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	//p.name = "Miss MoneyPenny" //works too
	(*p).name = "Miss MoneyPenny"

}

func main() {

	p1 := person{
		name: "Steve",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
