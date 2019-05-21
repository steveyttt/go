//This code will not work
//channels block!!!!
//CHANNELS BLOCK!!!!
//CHANNELS BLOCK!!!!
//CHANNELS BLOCK!!!!
//channels block!!!!
//you cannot just slam somthing on a channel and pull it off
//you need to orchestrate putting and pullng values off a channel in certain way
//think of it like a baton in a relay race
//you can only pass and accept when two things meet
package main

import "fmt"

func main() {
	c := make(chan int)

	//a new go routine is created and this CAN add to the c channel
	go func() {
		c <- 42
	}()

	//channel BLOCK - so fmt.println(<-c) will not print untuil the aboev func puts something on the channel
	fmt.Println(<-c)
}
