//https://golang.org/ref/spec#String_types
//https://blog.golang.org/strings
package main

import "fmt"

func main() {
	s := `"Hello,
	  SteveyT"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	//This prints the byte make up from s
	//it will print the literal bytes that make up the string Hello SteveyT
	//you can check them bytes inside the ASCII wiki page
	//Cross reference the decimal numbers against the letter
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	//Print hexadecimal
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		// fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
