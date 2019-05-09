package main

import (
	"fmt"
	"sort"
)

//Person to go outside
type Person struct {
	First string
	Age   int
}

//ByName to go to outside
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {

	p1 := Person{"James", 32}
	p2 := Person{"MoneyPenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)

}
