package main

import "fmt"

//a returned function is one which can present variables as a variable back
//func (r receiver) identifier (parameters) (return (s)) {function code}

func main() {
	foo()          //func with no returns
	bar("StevenT") //func with no returns

	s1 := woo("MoneyPenny") //function which can return values (store as variables)
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleming") //function which takes 2 parameters / arguments
	fmt.Println(x)
	fmt.Println(y)
}

//basic function
func foo() {
	fmt.Println("hello there matey moo")
}

//func bar expects a variable which we call s which is of type string
func bar(s string) {
	fmt.Println("Hello", s)
}

//func woo expects one parameter we call s and MUST be of type string
//it will the RETURN a string (The second reference to string)
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

//func mouse takes 2 string arguments and returns a string and a bool
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "hello"`)
	b := false
	return a, b
}
