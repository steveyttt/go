package bob

import "fmt"

//Testme is a small test to show how to export a func
func Testme(x string) {
	fmt.Println(x, "is your string")
}

func dontTestme(x string) {
	fmt.Println(x, "is your string")
}
