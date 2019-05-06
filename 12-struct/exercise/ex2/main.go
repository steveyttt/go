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
		last:        "BoyTyson",
		favicecream: []string{"chocolate", "pistachio"},
	}

	p2 := person{
		first:       "Grace",
		last:        "GirlTyson",
		favicecream: []string{"vanilla", "bubblegum"},
	}

	//create a map holding different structure types
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println("new human")
		fmt.Println("Key is", k, "value is", v)
		fmt.Println("first name is", v.first)
		for _, v := range v.favicecream {
			fmt.Println("Fav icecream is", v)
		}
	}

	fmt.Println(m)

}
