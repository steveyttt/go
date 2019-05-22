package main

import (
	"fmt"
	"log"
	"os"
)

//fatal error - it's dead stop immediately
//panic error - shits gotten weird, run defered functions and then try to exit UNLESS recover error
//--------defered functions run in the order of last in first out when panicking
//recover error - can recover a PANIC scenario, not always a good idea though.
//--------you can only run recover inside a defered function

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("LOG ----- ERROR OCCURED", err)
		fmt.Println("FMT ----- ERROR OCCURED", err)
		log.Fatal("FATAL ----- ERROR OCCURED", err)
		log.Panic("PANIC ----- ERROR OCCURED", err)
		panic(err)

	}
}
