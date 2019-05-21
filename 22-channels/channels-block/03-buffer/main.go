//this does work as you have set the buffer in the channel (1)
//by setting this you can simply pull one value off the channel
//This is hard to programme with though, you need to know exavctly how many values will be added to your channel
//buffer channels allow one value to sit in a channel regardless of whether something is ready to pull if off
package main

import "fmt"

func main() {
	c := make(chan []string, 1)
	c <- []string{"boo", "you", "too"}
	fmt.Println(<-c)
}
