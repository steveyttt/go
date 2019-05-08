package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	//working matey

	s1 := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"MoneyPenny","Age":27}]`
	bs := []byte(s1)

	people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Printf("%T\n", bs)
	fmt.Printf("%T\n", s1)

}
