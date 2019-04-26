package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp)

	//create a slice within a slice (nested)
	xpsd := [][]string{jb, mp}
	fmt.Println(xpsd)

	xp := append(jb, mp...)
	fmt.Println(xp)

}
