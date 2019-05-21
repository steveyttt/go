package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Sending", i, "to the channel")
			c <- i
		}
		close(c) //you need to close a channel when youre done with it - you only need clsoe whn using a for loop
	}()

	//receive
	for v := range c {
		fmt.Println("Receiving", v, "from the channel")
	}
	fmt.Println("about to exit")

}
