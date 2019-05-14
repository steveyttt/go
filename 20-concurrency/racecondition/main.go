package main

//go routines fighting over shared vars

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched() //releases the running routine to start the next (Not necessarily needed, but handy to know)
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Number of go routines", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Number of go routines", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
