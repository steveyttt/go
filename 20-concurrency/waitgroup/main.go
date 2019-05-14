package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) //wait for one additional go routine to complete

	go foo() //stick go before it to make it run inside it's own go routine
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait() //wait until all waitgroups have completed before moving to next step

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() //send a done message when the function completes. (You can see above wg.Add(1) is waiting for one DONE message)
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
