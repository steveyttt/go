package main

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
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Number of go routines", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Number of go routines", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
