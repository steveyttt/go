package main

//use the atomic pacage to write to a shared var inside a series of go routines
//try and avoid this though, stick with channels

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)                         //increase the counter value by 1
			fmt.Println("counter\t", atomic.LoadInt64(&counter)) //read counter
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Number of go routines", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Number of go routines", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
