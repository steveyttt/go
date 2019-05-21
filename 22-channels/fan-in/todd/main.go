package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v, "ranging over fan in")
		// fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			fmt.Println(i, "sending to channel even")
			even <- i
		} else {
			fmt.Println(i, "sending to channel odd")
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fmt.Println(v, "fan in to channel fan in")
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fmt.Println(v, "fan in to channel fan in")
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
