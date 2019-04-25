package main

import "fmt"

func main() {
	//for loop
	//for init; condition; post;
	//for init(initial start); condition(if true run again); post(after each run do this);
	//https://golang.org/ref/spec#IncDec_statements
	//for statement with for clause
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

}
