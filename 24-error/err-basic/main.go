package main

import "fmt"

func main() {
	n, err := fmt.Println("hello")
	//chek your errors rigth after running something!!!
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
