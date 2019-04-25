package main

import "fmt"

func main() {
	//switch evaluates a conditional statement and executes on first match only
	//if you add the FALLTHROUGH keyword it will allow all other matching case statements below to also execute
	//avoid FALLTHROUGH though - it drops wonky logic everywhere.
	//switch on a true boolean
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print2")
	case (3 == 3):
		fmt.Println("This should print cos 3 == 3")
		fallthrough
	case (4 == 4):
		fmt.Println("This should now print")
	default: // executed if no statements match
		fmt.Println("This is default") //if nothing matches run the DEFAULT
	}

	//switch on a value
	//if bond is matched in any case then execute
	switch "Bond" {
	case "MoneyPenny", "Bond", "Frank":
		fmt.Println("I am the money penny or bond or Frank")
	case "Q":
		fmt.Println("I am the Q")
	}
}
