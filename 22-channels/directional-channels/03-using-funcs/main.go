package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go foo(c)

	//receive
	//do not run this concurrently by prepending with go
	//this will naturally wait until the above func completes before running
	bar(c)

	fmt.Println("about to exit")

}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
