package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("names.txt") //open names.txt
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() //defer closure until the end

	bs, err := ioutil.ReadAll(f) //read all content in names.txt
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs)) //print bs as a string

}
