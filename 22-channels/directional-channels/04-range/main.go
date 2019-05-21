package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	for v := range c {
		fmt.Println("Receiving", v, "from the channel")
	}
	fmt.Println("about to exit")

}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Sending", i, "to the channel")
		c <- i
	}
	close(c) //you need to close a channel when youre done with it - if you dont close it you cannot take values off it as it is waiting on more values to come
}
