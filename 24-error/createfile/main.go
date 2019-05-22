package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//create a file called names.txt
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	//close the file when the function exits (defer)
	defer f.Close()

	//put whats up in buffer
	r := strings.NewReader("Wassup")

	//copies wassup (r) into names.txt (f)
	//it's file then value order
	io.Copy(f, r)

}
