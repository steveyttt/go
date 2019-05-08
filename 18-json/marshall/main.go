package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "MoneyPenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	var out bytes.Buffer
	json.Indent(&out, bs, "=", "\t")
	out.WriteTo(os.Stdout)

}
