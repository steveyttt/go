package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("About to exit")

}
