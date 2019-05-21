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

	//put 42 onto the channel
	c <- 42

	//take 42 off the channel
	fmt.Println(<-c)
}
