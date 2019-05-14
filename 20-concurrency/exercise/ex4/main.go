package main

//mutex - lock a shared variable that all go routines will use.
//it's like checking a book out in a library, once someone has it they need to return it to allow someone else to check it out

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

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() //this locks all variables accessed inside the go routine
			v := counter
			v++
			counter = v
			mu.Unlock() //this unlocks variables accessed inside the go routine
			wg.Done()
		}()
		fmt.Println("Number of go routines", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("Number of go routines", runtime.NumGoroutine())
	fmt.Println("count:", counter)

}
