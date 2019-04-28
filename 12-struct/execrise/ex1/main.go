package main

import "fmt"

type person struct {
	first       string
	last        string
	favicecream []string
}

func main() {

	p1 := person{
		first:       "George",
		last:        "Tyson",
		favicecream: []string{"chocolate", "pistachio"},
	}

	p2 := person{
		first:       "Grace",
		last:        "Tyson",
		favicecream: []string{"vanilla", "bubblegum"},
	}

	fmt.Println(p1.favicecream)
	fmt.Println(p2.favicecream)

	for k, v := range p1.favicecream {
		fmt.Println(k, v)
	}

	for k, v := range p2.favicecream {
		fmt.Println(k, v)
	}

}
