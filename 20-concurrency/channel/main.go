package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	//make a channel
	ch := make(chan int)
	go func() {
		//put a value on the channel
		ch <- doSomething(5)
	}()
	//print out the value on the channel
	fmt.Println(<-ch)

	go func() {
		ch <- doSomething(6)
	}()
	fmt.Println(<-ch)

}
